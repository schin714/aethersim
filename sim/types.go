package sim

// Character holds all mutable state for a combatant.
type Character struct {
	ID        string
	IsAlly    bool
	Level     int
	Health    float64
	MaxHealth float64
	Mana      float64
	MaxMana   float64
	Strength  float64
	Defense   float64
	Spirit    float64
	Speed     float64
	Evasion   float64
	Shields   int
	Elements  []Element // e.g. Fire, Water, Earth, etc.
	Abilities []*Ability

	ActiveBuffs   []BuffInstance
	ActiveDebuffs []DebuffInstance
}

// BuffInstance represents one application of a buff.
type BuffInstance struct {
	AppliedBy     string  // who applied this buff
	ModifierPct   float64 // +25 for +25% buff
	TotalRounds   int     // number of turns remaining
	RoundsApplied int     // how many rounds have ticked
	Stat          string  // e.g. "Strength" or "Defense"
}

// DebuffInstance represents one application of a debuff.
type DebuffInstance struct {
	AppliedBy      string
	ModifierPct    float64 // –25 for –25% defense-down, or 0 if it's a pure DoT
	DamagePercent  float64 // >0 if it deals DoT each round
	ElementToApply Element // element to apply as debuff, if any
	TotalRounds    int
	RoundsApplied  int
	Stat           string  // e.g. "Defense" or empty if pure DoT
	Element        Element // for elemental modifiers, if needed
}
