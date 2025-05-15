package sim

// Buff represents a temporary, positive modifier.
type Buff struct {
	Type            string  // "defense", "strength", etc.
	Rounds          int     // number of turns it lasts
	ModifierPercent float64 // e.g. +25 for +25%
}

// Debuff represents a temporary, negative effect (or DoT).
type Debuff struct {
	Type              string  // "poison", "burn", "defense", etc.
	Element           Element // element for elemental debuffs
	Rounds            int     // duration in turns
	DamagePercent     float64 // percentage of max HP per tick, if any
	ModifierPercent   float64 // e.g. –25 for –25% reduction
	ApplicationChance float64 // percent chance to land
	ElementToApply    Element // element to apply as debuff, if any
}

// Ability is the core data you need for simulation.
type Ability struct {
	ID               string  // unique key
	Power            float64 // base dmg or heal amount
	Buff             *Buff   // non-nil if this is a buff ability
	Debuff           *Debuff // non-nil if this is a debuff ability
	Type             string  // "attack", "heal", "buff", "debuff"
	TargetType       string  // "single", "all", "self"
	TargetSelectType string  // "ally", "enemy", "any"
	ManaCost         float64 // e.g. 2 or –1 if none
	Element          Element // elemental affiliation
}

// AbilityDict lets you look up an Ability by its key.
var AbilityDict = map[string]*Ability{
	"bash": {
		ID:               "bash",
		Power:            30,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          "wild",
	},
	"fierce-bite": {
		ID:               "fierce-bite",
		Power:            40,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          "wild",
	},
	"leaf-slash": {
		ID:               "leaf-slash",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          "earth",
	},
	"nature-blessing": {
		ID:               "nature-blessing",
		Power:            50,
		Type:             "heal",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         2,
		Element:          "earth",
	},
	"earthquake": {
		ID:               "earthquake",
		Power:            70,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         3,
		Element:          "earth",
	},
	"barkskin": {
		ID:               "barkskin",
		Power:            0,
		Type:             "buff",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         1,
		Element:          "earth",
		Buff: &Buff{
			Type:            "defense",
			Rounds:          4,
			ModifierPercent: 25,
		},
	},
	"poison-spray": {
		ID:               "poison-spray",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         1,
		Element:          "earth",
		Debuff: &Debuff{
			Type:              "poison",
			Element:           "earth",
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 100,
		},
	},
	"ice-shard": {
		ID:               "ice-shard",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Ice,
	},
	"blizzard": {
		ID:               "blizzard",
		Power:            50,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Ice,
	},

	// Water
	"squirt": {
		ID:               "squirt",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Water,
	},
	"soak": {
		ID:               "soak",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "any",
		ManaCost:         1,
		Element:          Water,
		Debuff: &Debuff{
			Type:              "element",
			Element:           Water,
			Rounds:            4,
			ApplicationChance: 100,
			ElementToApply:    Water,
		},
	},
	"tidal-wave": {
		ID:               "tidal-wave",
		Power:            50,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Water,
	},

	// Fire
	"flame-burp": {
		ID:               "flame-burp",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Fire,
	},
	"flame-lash": {
		ID:               "flame-lash",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Fire,
	},
	"molten-burst": {
		ID:               "molten-burst",
		Power:            60,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Fire,
	},
	"scorch": {
		ID:               "scorch",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         1,
		Element:          Fire,
		Debuff: &Debuff{
			Type:              "burn",
			Element:           Fire,
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 100,
		},
	},

	// Electric
	"lightning-bolt": {
		ID:               "lightning-bolt",
		Power:            40,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Electric,
	},
	"lightning-storm": {
		ID:               "lightning-storm",
		Power:            70,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         3,
		Element:          Electric,
	},

	// Astral
	"celestial-beam": {
		ID:               "celestial-beam",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Astral,
	},
	"astral-veil": {
		ID:               "astral-veil",
		Power:            0,
		Type:             "buff",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         2,
		Element:          Astral,
		Buff: &Buff{
			Type:            "shield",
			Rounds:          1,
			ModifierPercent: 2,
		},
	},

	// Character-specific...
	"chocolate-blast": {
		ID:               "chocolate-blast",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
	},
	"sweet-heal": {
		ID:               "sweet-heal",
		Power:            40,
		Type:             "heal",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         2,
		Element:          Earth,
	},
	"cookie-crunch": {
		ID:               "cookie-crunch",
		Power:            50,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         0,
		Element:          Wild,
	},

	// Wolfpup
	"inspiring-howl": {
		ID:               "inspiring-howl",
		Power:            50,
		Type:             "heal",
		TargetType:       "all",
		TargetSelectType: "ally",
		ManaCost:         2,
		Element:          Earth,
	},

	// Verdant Viper
	"viper-strike": {
		ID:               "viper-strike",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
	},

	// Sporepuff
	"spore-shot": {
		ID:               "spore-shot",
		Power:            25,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "poison",
			Element:           Earth,
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 50,
		},
	},
	"poison-spores": {
		ID:               "poison-spores",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         1,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "poison",
			Element:           Earth,
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 100,
		},
	},
	"rotting-grasp": {
		ID:               "rotting-grasp",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         1,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "defense",
			Element:           Earth,
			Rounds:            4,
			ModifierPercent:   20,
			ApplicationChance: 100,
		},
	},
	"toxic-cloud": {
		ID:               "toxic-cloud",
		Power:            20,
		Type:             "debuff",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         3,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "poison",
			Element:           Earth,
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 75,
		},
	},

	// Breezeling
	"gale-cut": {
		ID:               "gale-cut",
		Power:            35,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Air,
	},
	"updraft": {
		ID:               "updraft",
		Power:            0,
		Type:             "buff",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         1,
		Element:          Air,
		Buff: &Buff{
			Type:            "evasion",
			Rounds:          5,
			ModifierPercent: 15,
		},
	},
	"whirling-feathers": {
		ID:               "whirling-feathers",
		Power:            50,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Air,
	},

	// Stonebound Sentinel
	"boulder-bash": {
		ID:               "boulder-bash",
		Power:            40,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
	},
	"iron-bulwark": {
		ID:               "iron-bulwark",
		Power:            0,
		Type:             "buff",
		TargetType:       "single",
		TargetSelectType: "ally",
		ManaCost:         1,
		Element:          Earth,
		Buff: &Buff{
			Type:            "defense",
			Rounds:          4,
			ModifierPercent: 30,
		},
	},
	"earthen-grasp": {
		ID:               "earthen-grasp",
		Power:            0,
		Type:             "debuff",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         1,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "defense",
			Element:           Earth,
			Rounds:            4,
			ModifierPercent:   25,
			ApplicationChance: 80,
		},
	},
	"stone-spire": {
		ID:               "stone-spire",
		Power:            60,
		Type:             "attack",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Earth,
	},
	"decay-claw": {
		ID:               "decay-claw",
		Power:            30,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
	},
	"rot-tide": {
		ID:               "rot-tide",
		Power:            20,
		Type:             "debuff",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "poison",
			Element:           Earth,
			Rounds:            4,
			DamagePercent:     6,
			ApplicationChance: 100,
		},
	},
	"bone-spewer": {
		ID:               "bone-spewer",
		Power:            25,
		Type:             "attack",
		TargetType:       "single",
		TargetSelectType: "enemy",
		ManaCost:         -1,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "defense",
			Element:           Earth,
			Rounds:            3,
			ModifierPercent:   20,
			ApplicationChance: 100,
		},
	},
	"oak-shackle": {
		ID:               "oak-shackle",
		Power:            0,
		Type:             "debuff",
		TargetType:       "all",
		TargetSelectType: "enemy",
		ManaCost:         2,
		Element:          Earth,
		Debuff: &Debuff{
			Type:              "strength",
			Element:           Earth,
			Rounds:            3,
			ModifierPercent:   20,
			ApplicationChance: 100,
		},
	},
}
