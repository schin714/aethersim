package sim

import "math/rand"

// RandomDecision picks a random usable ability and then selects valid targets.
func RandomDecision(
	actor *Character,
	allies, enemies []*Character,
) (*Ability, []*Character) {
	// 1) filter abilities by mana
	usable := make([]*Ability, 0, len(actor.Abilities))
	for _, ab := range actor.Abilities {
		if actor.Mana >= ab.ManaCost {
			usable = append(usable, ab)
		}
	}
	if len(usable) == 0 {
		return nil, nil
	}

	// 2) pick one at random
	ab := usable[rand.Intn(len(usable))]

	// 3) select targets based on ab.TargetType / ab.TargetSelectType
	var pool []*Character
	switch ab.TargetType {
	case "self":
		return ab, []*Character{actor}
	case "single":
		if ab.TargetSelectType == "ally" {
			if actor.IsAlly {
				pool = allies
			} else {
				pool = enemies
			}
		} else if ab.TargetSelectType == "enemy" {
			if actor.IsAlly {
				pool = enemies
			} else {
				pool = allies
			}
		} else {
			pool = append(allies, enemies...)
		}
		if len(pool) == 0 {
			// no valid target: just return nil or skip this ability
			return nil, nil
		}
		// only one target
		return ab, []*Character{pool[rand.Intn(len(pool))]}
	case "all":
		if ab.TargetSelectType == "ally" {
			if actor.IsAlly {
				return ab, allies
			} else {
				return ab, enemies
			}
		} else if ab.TargetSelectType == "enemy" {
			if actor.IsAlly {
				return ab, enemies
			} else {
				return ab, allies
			}
		} else {
			return ab, append(allies, enemies...)
		}
	}

	// fallback: no valid targets
	return ab, nil
}
