package sim

import (
	"math"
	"math/rand"
)

// UtilityDecision scores every (ability, target) combo and does a roulette‐wheel pick.
func UtilityDecision(
	actor *Character,
	allies, enemies []*Character,
) (*Ability, []*Character) {
	// 1) filter by usable mana
	var usable []*Ability
	for i := range actor.Abilities {
		ab := actor.Abilities[i]
		if actor.Mana >= ab.ManaCost {
			usable = append(usable, ab)
		}
	}
	if len(usable) == 0 {
		return nil, nil
	}

	// 1a) low-health heal
	if actor.Health < 0.3*actor.MaxHealth {
		for _, ab := range usable {
			if ab.Type == "heal" {
				return ab, []*Character{actor}
			}
		}
	}

	// 2) build combos
	type combo struct {
		ab    *Ability
		tgt   *Character
		score float64
	}
	var combos []combo

	for _, ab := range usable {
		// pick pool
		var pool []*Character
		if ab.TargetSelectType == "ally" {
			if actor.IsAlly {
				pool = allies
			} else {
				pool = enemies
			}

		} else {
			if actor.IsAlly {
				pool = enemies
			} else {
				pool = allies
			}
		}
		for _, tgt := range pool {
			if tgt.Health <= 0 {
				continue
			}
			// skip overheal
			if ab.Type == "heal" && tgt.Health >= tgt.MaxHealth {
				continue
			}
			s := scoreCombo(actor, ab, tgt, allies, enemies)
			combos = append(combos, combo{ab, tgt, s})
		}
	}
	if len(combos) == 0 {
		return nil, nil
	}

	// 3) roulette‐wheel
	total := 0.0
	for _, c := range combos {
		if c.score > 0 {
			total += c.score
		}
	}
	if total > 0 {
		roll := rand.Float64() * total
		for _, c := range combos {
			if c.score <= 0 {
				continue
			}
			roll -= c.score
			if roll <= 0 {
				return c.ab, []*Character{c.tgt}
			}
		}
	}

	// 4) fallback: best‐score
	best := combos[0]
	for _, c := range combos[1:] {
		if c.score > best.score {
			best = c
		}
	}
	return best.ab, []*Character{best.tgt}
}

// scoreCombo is the direct translation of your evaluateCombo logic.
func scoreCombo(
	src *Character,
	ab *Ability,
	tgt *Character,
	allies, enemies []*Character,
) float64 {
	score := 0.0

	switch ab.Type {
	case "heal":
		missing := tgt.MaxHealth - tgt.Health
		score += float64(ab.Power) * (1 + missing/tgt.MaxHealth)

	case "buff":
		if ab.Buff.Type == "shield" {
			score += ab.Buff.ModifierPercent * 15
		} else {
			score += ab.Buff.ModifierPercent * 0.3 * float64(ab.Buff.Rounds)
		}
		score += 10 // straight buff bonus

	case "debuff":
		d := ab.Debuff
		if d.DamagePercent > 0 {
			score += (d.ApplicationChance / 100.0) * d.DamagePercent * float64(d.Rounds)
		} else if d.ModifierPercent > 0 {
			score += (d.ApplicationChance / 100.0) * (d.ModifierPercent * 0.3) * float64(d.Rounds)
		}
		score += 10 // straight debuff bonus

	default: // Attack
		base := float64(ab.Power)
		bonus := 0.0
		if tgt.Health/tgt.MaxHealth < 0.3 {
			bonus = 0.5
		}
		score += base * (1 + bonus)
		if base*(tgt.Health/200) >= tgt.Health {
			score += 20
		}
	}

	// AOE
	if ab.TargetType == "all" {
		var pool []*Character
		if ab.TargetSelectType == "ally" {
			pool = allies
		} else {
			pool = enemies
		}
		if len(pool) > 1 {
			score *= 1.2
		} else {
			score *= 0.1
		}
	}

	// element
	if ab.Type == "attack" {
		weak, resist := false, false
		for _, el := range tgt.Elements {
			for _, w := range GetWeaknesses(el) {
				if w == ab.Element {
					weak = true
				}
			}
			for _, r := range GetResistances(el) {
				if r == ab.Element {
					resist = true
				}
			}
		}
		if weak {
			score *= 2
		}
		if resist {
			score *= 0.5
		}
	}

	// mana
	mr := src.Mana / src.MaxMana
	scarcity := 1 - mr
	costNorm := float64(ab.ManaCost) / src.MaxMana
	const (
		costWeight        = 1.0
		genScarcityWeight = 1.0
		genCritWeight     = 1.0
		fullManaBonus     = 0.6
	)
	if ab.ManaCost > 0 {
		score *= math.Max(0, 1-costNorm*scarcity*costWeight)
	} else if ab.ManaCost < 0 {
		inner := scarcity*genScarcityWeight + MAX_MANA_CRIT_CHANCE*genCritWeight
		if mr == 1 {
			inner += fullManaBonus
		}
		score *= 1 + -costNorm*inner
	}

	// repeat‐debuff
	if ab.Type == "debuff" {
		for _, inst := range tgt.ActiveDebuffs {
			if inst.Stat == ab.Debuff.Type {
				if inst.Stat == "element" {
					score -= 50
				} else {
					score -= 10 * float64(inst.TotalRounds-inst.RoundsApplied)
				}
			}
		}
	}

	// randomness
	score += rand.Float64() * 0.1

	return score
}
