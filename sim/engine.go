package sim

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
)

type Engine struct {
	Characters  []*Character
	TurnOrder   []int // indices into Characters
	Current     int   // index into TurnOrder
	PlayerWon   bool
	GameOver    bool
	LastImpacts []ImpactRecord // last round’s impacts, for stats
	TotalRounds int
	TotalTurns  int
}

type ImpactRecord struct {
	ActorID  string  // who caused it
	TargetID string  // who received it
	Delta    float64 // positive = damage, negative = heal
	IsDebuff bool    // true if it was a debuff (e.g. poison, burn, etc.)
}

const ELEMENTAL_EFFECTIVENESS_MODIFIER = 1.5

// NewEngine constructs a fresh engine from two teams.
func NewEngine(allies, enemies []*Character) *Engine {
	chars := append(allies, enemies...)
	n := len(chars)

	// create a slice of indices [0,1,2...n-1]
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}

	// sort descending by chars[index].Speed, tiebreaker random
	sort.Slice(order, func(i, j int) bool {
		ci, cj := chars[order[i]], chars[order[j]]
		if ci.Speed != cj.Speed {
			return ci.Speed > cj.Speed
		}
		// same Speed → flip a coin
		return rand.Intn(2) == 0
	})

	return &Engine{
		Characters:  chars,
		TurnOrder:   order,
		GameOver:    false,
		PlayerWon:   false,
		Current:     0,
		TotalRounds: 0,
		TotalTurns:  0,
	}
}

// Reset brings the engine back to round 1 with fresh stats.
func (e *Engine) Reset(allies, enemies []*Character) {
	*e = *NewEngine(allies, enemies)
}

// one turn’s logic: choose ability & target, apply effects
func (e *Engine) Step(decisionFn func(actor *Character, allies, enemies []*Character) (*Ability, []*Character)) {
	e.LastImpacts = e.LastImpacts[:0] // reset slice, reuse capacity
	e.TotalTurns++
	e.checkEnd()
	if e.GameOver {
		return
	}
	actorIdx := e.TurnOrder[e.Current]
	actor := e.Characters[actorIdx]
	if actor.Health <= 0 {
		e.advanceTurn()
		return
	}

	// Partition alive allies and enemies
	var allies, enemies []*Character
	for _, c := range e.Characters {
		if c.Health <= 0 {
			continue
		}
		if c.IsAlly {
			allies = append(allies, c)
		} else {
			enemies = append(enemies, c)
		}
	}

	// Decision: pick ability AND targets
	ability, targets := decisionFn(actor, allies, enemies)
	if ability == nil || len(targets) == 0 {
		// no valid action—just skip turn
		fmt.Printf("DEBUG: %s has no valid action\n", actor.ID)
		e.advanceTurn()
		return
	}

	// Pay mana
	actor.Mana = clamp(actor.Mana-ability.ManaCost, 0, actor.MaxMana)

	// apply effects
	for _, tgt := range targets {
		delta := e.applyAbilityToTarget(actor, tgt, ability)
		// fmt.Printf("DEBUG: %s → %s = %0.1f\n", actor.ID, tgt.ID, delta)
		e.LastImpacts = append(e.LastImpacts, ImpactRecord{
			ActorID:  actor.ID,
			TargetID: tgt.ID,
			Delta:    delta,
			IsDebuff: false,
		})
	}

	// check win
	e.checkEnd()
	if !e.GameOver {
		e.advanceTurn()
	}
}

func (e *Engine) advanceTurn() {
	prev := e.Current
	e.Current = (e.Current + 1) % len(e.TurnOrder)
	if prev > e.Current {
		// just wrapped a full round
		e.HandleNewRoundEffects()
		e.TotalRounds++
	}
}

func (e *Engine) checkEnd() {
	alives := map[bool]bool{true: false, false: false}
	for _, c := range e.Characters {
		if c.Health > 0 {
			alives[c.IsAlly] = true
		}
	}
	if !alives[true] || !alives[false] || e.TotalRounds >= 100 {
		e.GameOver = true
		e.PlayerWon = alives[true]
	}
}

// HandleNewRoundEffects ticks every debuff for DoT and prunes expired effects.
// Call it once per round; it mutates character health and effect lists in place.
func (e *Engine) HandleNewRoundEffects() {
	for _, c := range e.Characters {
		if c.Health <= 0 {
			continue
		}

		// 1) Deal DoT from any debuffs that have a DamagePercent > 0
		totalDot := 0.0
		for i := range c.ActiveDebuffs {
			db := &c.ActiveDebuffs[i]
			if db.DamagePercent > 0 {
				elementalMod := 1.0
				for _, elem := range c.Elements {
					for _, w := range GetWeaknesses(elem) {
						if w == db.Element {
							elementalMod *= ELEMENTAL_EFFECTIVENESS_MODIFIER
						}
					}
					for _, s := range GetResistances(elem) {
						if s == db.Element {
							elementalMod /= ELEMENTAL_EFFECTIVENESS_MODIFIER
						}
					}
				}
				dot := math.Ceil(c.MaxHealth * (db.DamagePercent / 100) * elementalMod)
				totalDot += dot

				e.LastImpacts = append(e.LastImpacts, ImpactRecord{
					ActorID:  db.AppliedBy,
					TargetID: c.ID,
					Delta:    dot,
					IsDebuff: true,
				})
			}
		}
		if totalDot > 0 {
			c.Health = math.Max(0, c.Health-totalDot)
		}

		// 2) Expire buffs (they only modified damage when you cast; no stat rollback needed)
		for i := len(c.ActiveBuffs) - 1; i >= 0; i-- {
			b := &c.ActiveBuffs[i]
			b.RoundsApplied++
			if b.RoundsApplied >= b.TotalRounds {
				// just remove it
				c.ActiveBuffs = slices.Delete(c.ActiveBuffs, i, i+1)
			}
		}

		// 3) Expire debuffs (non‐DoT side effects—elements, stat‐mods—are only used at resolution)
		for i := len(c.ActiveDebuffs) - 1; i >= 0; i-- {
			db := &c.ActiveDebuffs[i]
			db.RoundsApplied++
			if db.RoundsApplied >= db.TotalRounds {
				c.ActiveDebuffs = slices.Delete(c.ActiveDebuffs, i, i+1)
			}
		}
	}
}

// applyAbilityToTarget applies one Ability cast by `source` onto `target`.
// It mutates target.Health and pushes any new BuffInstance or DebuffInstance
// onto the target or source as appropriate.
// Returns the raw impact (positive = damage dealt, negative = healing done).
func (e *Engine) applyAbilityToTarget(
	source *Character,
	target *Character,
	ability *Ability,
) float64 {
	// 1) calculate base impact: damage or heal
	impact := 0.0
	switch ability.Type {
	case "attack", "debuff":
		if rand.Float64() < EffectiveEvasion(target) {
			return 0 // missed entirely
		}
		if target.Shields > 0 {
			target.Shields -= 1
			return 0
		}
		elementalMod := 1.0
		for _, elem := range target.Elements {
			for _, w := range GetWeaknesses(elem) {
				if w == ability.Element {
					elementalMod *= ELEMENTAL_EFFECTIVENESS_MODIFIER
				}
			}
			for _, s := range GetResistances(elem) {
				if s == ability.Element {
					elementalMod /= ELEMENTAL_EFFECTIVENESS_MODIFIER
				}
			}
		}
		// apply strength buffs/debuffs on source
		strMod := GetEffectiveModifierForAttack(source, "strength")
		// apply defense buffs/debuffs on target
		defMod := GetEffectiveModifierForAttack(target, "defense")

		// basic damage formula
		base := ability.Power * float64(source.Strength+5) / math.Pow(float64(target.Defense+13), 0.9) * (float64((source.Level*2)/4+5) / 30)
		aoeMod := 1.0
		if ability.TargetType == "all" {
			aoeMod = 0.67
		}
		damage := base * strMod * defMod * elementalMod * aoeMod
		impact = math.Ceil(damage)

	case "heal", "buff":
		// simple healing formula
		base := ability.Power * (source.Spirit / (17 + 14*math.Log(float64(source.Level)))) * HealMultiplier(e.TotalRounds)
		impact = math.Ceil(base)
	}

	// 2) apply to target health (or healing)
	if ability.Type == "heal" {
		target.Health = math.Min(target.MaxHealth, target.Health+impact)
	} else {
		target.Health = math.Max(0, target.Health-impact)
	}

	// 3) if ability grants a buff to the target (or source for self‐buff)
	if ability.Buff != nil {
		b := ability.Buff

		// 4a) shield type
		if b.Type == "shield" {
			// ModifierPct is the shield amount
			target.Shields += int(b.ModifierPercent)

		} else {
			// 4b) all other buffs: compute spirit‐scaled percentage
			adj := b.ModifierPercent
			if b.ModifierPercent != 0 {
				adj = b.ModifierPercent + (b.ModifierPercent/2)*(source.Spirit/200)
			}
			// push the new buff instance
			target.ActiveBuffs = append(target.ActiveBuffs, BuffInstance{
				Stat:          b.Type, // e.g. "defense" or "strength"
				ModifierPct:   adj,
				TotalRounds:   b.Rounds,
				RoundsApplied: 0,
			})
		}
	}

	// 4) if ability has a debuff, maybe apply it
	if ability.Debuff != nil {
		// 4a) roll for applicationChance
		if rand.Float64()*100 < ability.Debuff.ApplicationChance {
			db := ability.Debuff // alias for brevity
			// 4b) element‐type debuff: reset or add element
			if db.Type == "element" && db.ElementToApply != "" {
				// see if there’s an existing element‐debuff of this type
				for i := range target.ActiveDebuffs {
					inst := &target.ActiveDebuffs[i]
					if inst.Stat == db.Type {
						// reset its duration
						inst.RoundsApplied = 0
						goto appliedDone
					}
				}
				// if target already has the base element, skip
				for _, el := range target.Elements {
					if el == db.ElementToApply {
						goto appliedDone
					}
				}
				// otherwise add the element status
				target.Elements = append(target.Elements, db.ElementToApply)
			}

			// 4c) non‐element debuff (or element with no ElementToApply)
			// compute spirit‐scaled modifierPct if present
			adj := db.ModifierPercent
			if db.ModifierPercent != 0 {
				adj = db.ModifierPercent + (db.ModifierPercent/2)*(source.Spirit/200)
			}

			// push a fresh instance
			target.ActiveDebuffs = append(target.ActiveDebuffs, DebuffInstance{
				AppliedBy:      source.ID,
				Stat:           db.Type,
				ModifierPct:    adj,
				DamagePercent:  db.DamagePercent,
				ElementToApply: db.ElementToApply,
				TotalRounds:    db.Rounds,
				RoundsApplied:  0,
				Element:        db.Element,
			})
		}
	appliedDone:
		// regardless of success or skip, we keep the impact from the damage above
	}

	// return the raw amount of HP change (positive = damage; negative = heal)
	if ability.Type == "heal" || ability.Type == "buff" {
		return -impact
	}
	return impact
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
