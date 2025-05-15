package sim

import "math"

const MAX_MANA_CRIT_CHANCE = 0.34

// EffectiveEvasion returns the miss‐chance [0,0.75] for a character,
// combining base Evasion plus any “evasion” buffs.
func EffectiveEvasion(c *Character) float64 {
	base := c.Evasion
	totalBuffPct := 0.0
	for _, b := range c.ActiveBuffs {
		if b.Stat == "evasion" {
			totalBuffPct += b.ModifierPct
		}
	}
	ev := base + totalBuffPct/100
	if ev > 0.75 {
		return 0.75
	}
	return ev
}

// GetEffectiveModifierForAttack returns the multiplicative damage modifier
// based on strength‐type buffs/debuffs on source, or defense‐type on target.
// stat should be either "strength" or "defense".
func GetEffectiveModifierForAttack(c *Character, stat string) float64 {
	// sum up all matching buff percentages
	buffTotal := 0.0
	for _, b := range c.ActiveBuffs {
		if b.Stat == stat {
			buffTotal += b.ModifierPct
		}
	}
	// sum up all matching debuff percentages
	debuffTotal := 0.0
	for _, d := range c.ActiveDebuffs {
		if d.Stat == stat {
			debuffTotal += d.ModifierPct
		}
	}

	// clamp the net modifier between -50 and +50
	diff := buffTotal - debuffTotal
	if diff > 50 {
		diff = 50
	} else if diff < -50 {
		diff = -50
	}

	// for defense we invert the sign so a +25% defense buff reduces damage, etc.
	if stat == "defense" {
		diff = -diff
	}

	// now build the multiplicative modifier
	if diff > 0 {
		return 1 + diff/100
	} else if diff < 0 {
		return 1 / (1 + math.Abs(diff)/100)
	}
	return 1
}

func HealMultiplier(round int) float64 {
	const start, end = 15, 50
	const minMult = 0.15

	switch {
	case round < start:
		return 1
	case round >= end:
		return minMult
	default:
		// how far we are between start and end, 0→1
		t := float64(round-start) / float64(end-start)
		// linearly interpolate from 1→minMult
		return 1*(1-t) + minMult*t
	}
}
