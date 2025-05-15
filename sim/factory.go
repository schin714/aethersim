package sim

import "fmt"

// MakeTeam looks up each key in CharacterTemplates, applies the growth formulas,
// and returns a slice of *Character ready for battle.
func MakeTeam(keys []string, level int, isAlly bool) []*Character {
    team := make([]*Character, 0, len(keys))
    for _, key := range keys {
        tmpl, ok := CharacterTemplates[key]
        if !ok {
            panic(fmt.Sprintf("unknown character key %q", key))
        }
        c := &Character{
            ID:        key,
            IsAlly:    isAlly,
            Level:     level,
            MaxHealth: tmpl.BaseHealth + tmpl.HPGrowth*float64(level),
            Health:    tmpl.BaseHealth + tmpl.HPGrowth*float64(level),
            MaxMana:   tmpl.BaseMana,
            Mana:      tmpl.BaseMana / 2,
            Strength:  tmpl.BaseStrength + tmpl.StrengthGrowth*float64(level),
            Defense:   tmpl.BaseDefense + tmpl.DefenseGrowth*float64(level),
            Spirit:    tmpl.BaseSpirit + tmpl.SpiritGrowth*float64(level),
            Speed:     tmpl.BaseSpeed + tmpl.SpeedGrowth*float64(level),
            Evasion:   tmpl.Evasion,
			Shields:   0,
        }
		// Create a shallow copy of the elements slice to avoid modifying the original template
		elems := make([]Element, len(tmpl.Elements))
		copy(elems, tmpl.Elements)
		c.Elements = elems

        // Filter abilities the character can actually use at this level
        for _, at := range tmpl.AbilityTemplates {
			if at.MinLv <= level {
			  if ab, ok := AbilityDict[at.Key]; ok {
				c.Abilities = append(c.Abilities, ab)
			  }
			}
		  }
        team = append(team, c)
    }
    return team
}