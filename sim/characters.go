package sim

import "sort"

// AbilityTemplate ties an ability key to the minimum level required.
type CharacterAbilityTemplate struct {
	Key   string
	MinLv int
}

// Template holds all of the static, per‐species data you need to spawn a Character.
type CharacterTemplate struct {
	Elements         []Element
	BaseHealth       float64
	BaseMana         float64
	BaseStrength     float64
	BaseDefense      float64
	BaseSpirit       float64
	BaseSpeed        float64
	HPGrowth         float64
	StrengthGrowth   float64
	DefenseGrowth    float64
	SpiritGrowth     float64
	SpeedGrowth      float64
	Evasion          float64
	AbilityTemplates []CharacterAbilityTemplate
}

// AllCharacterKeys returns a sorted list of all character IDs in the templates map.
func AllCharacterKeys() []string {
	keys := make([]string, 0, len(CharacterTemplates))
	for id := range CharacterTemplates {
		keys = append(keys, id)
	}
	sort.Strings(keys)
	return keys
}

// CharacterTemplates maps your “id” → Template.
var CharacterTemplates = map[string]CharacterTemplate{
	// —————————————————————————————————————————————
	// chocolate_chip
	// —————————————————————————————————————————————
	"chocolate_chip": {
		Elements:       []Element{Earth},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2.25,
		SpiritGrowth:   2.25,
		SpeedGrowth:    1.5,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "chocolate-blast", MinLv: 1},
			{Key: "sweet-heal", MinLv: 3},
			{Key: "cookie-crunch", MinLv: 4},
		},
	},
	// —————————————————————————————————————————————
	// daring_wolfpup
	// —————————————————————————————————————————————
	"daring_wolfpup": {
		Elements:       []Element{Wild},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    5,
		BaseSpirit:     4,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "fierce-bite", MinLv: 2},
			{Key: "inspiring-howl", MinLv: 4},
		},
	},
	// —————————————————————————————————————————————
	// sprigshell
	// —————————————————————————————————————————————
	"sprigshell": {
		Elements:       []Element{Earth},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    6,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       5,
		StrengthGrowth: 1.5,
		DefenseGrowth:  3,
		SpiritGrowth:   2,
		SpeedGrowth:    1.5,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "leaf-slash", MinLv: 1},
			{Key: "barkskin", MinLv: 2},
			{Key: "nature-blessing", MinLv: 4},
		},
	},
	// —————————————————————————————————————————————
	// frostnip
	// —————————————————————————————————————————————
	"frostnip": {
		Elements:       []Element{Ice},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2.25,
		DefenseGrowth:  2.25,
		SpiritGrowth:   1.5,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "ice-shard", MinLv: 2},
			{Key: "blizzard", MinLv: 4},
		},
	},

	// —————————————————————————————————————————————
	// lightning_kat
	// —————————————————————————————————————————————
	"lightning_kat": {
		Elements:       []Element{Electric},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "lightning-bolt", MinLv: 1},
			{Key: "lightning-storm", MinLv: 4},
		},
	},

	// —————————————————————————————————————————————
	// verdant_viper
	// —————————————————————————————————————————————
	"verdant_viper": {
		Elements:       []Element{Earth},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      4,
		HPGrowth:       3,
		StrengthGrowth: 3.5,
		DefenseGrowth:  1.25,
		SpiritGrowth:   1.25,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "viper-strike", MinLv: 2},
			{Key: "poison-spray", MinLv: 3},
		},
	},

	// —————————————————————————————————————————————
	// giant_capy
	// —————————————————————————————————————————————
	"giant_capy": {
		Elements:       []Element{Earth},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    6,
		BaseSpirit:     3,
		BaseSpeed:      5,
		HPGrowth:       5,
		StrengthGrowth: 2,
		DefenseGrowth:  3.5,
		SpiritGrowth:   1,
		SpeedGrowth:    1.5,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "chocolate-blast", MinLv: 0},
			{Key: "earthquake", MinLv: 4},
		},
	},

	// —————————————————————————————————————————————
	// flitterfyre
	// —————————————————————————————————————————————
	"flitterfyre": {
		Elements:       []Element{Fire},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "flame-burp", MinLv: 1},
			{Key: "scorch", MinLv: 3},
		},
	},

	// —————————————————————————————————————————————
	// pondril
	// —————————————————————————————————————————————
	"pondril": {
		Elements:       []Element{Water},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "squirt", MinLv: 0},
			{Key: "soak", MinLv: 2},
		},
	},

	// —————————————————————————————————————————————
	// sporepuff
	// —————————————————————————————————————————————
	"sporepuff": {
		Elements:       []Element{Earth},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   4,
		BaseDefense:    5,
		BaseSpirit:     6,
		BaseSpeed:      5,
		HPGrowth:       5,
		StrengthGrowth: 1.5,
		DefenseGrowth:  2,
		SpiritGrowth:   3,
		SpeedGrowth:    1.5,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "spore-shot", MinLv: 0},
			{Key: "poison-spores", MinLv: 1},
			{Key: "rotting-grasp", MinLv: 3},
			{Key: "toxic-cloud", MinLv: 6},
		},
	},

	// —————————————————————————————————————————————
	// breezeling
	// —————————————————————————————————————————————
	"breezeling": {
		Elements:       []Element{Air},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    4,
		BaseSpirit:     4,
		BaseSpeed:      7,
		HPGrowth:       3,
		StrengthGrowth: 3,
		DefenseGrowth:  1.25,
		SpiritGrowth:   1.25,
		SpeedGrowth:    2.5,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "gale-cut", MinLv: 0},
			{Key: "updraft", MinLv: 1},
			{Key: "whirling-feathers", MinLv: 3},
		},
	},
	// —————————————————————————————————————————————
	// cinder_chip
	// —————————————————————————————————————————————
	"cinder_chip": {
		Elements:       []Element{Fire},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    5,
		BaseSpirit:     4,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 3,
		DefenseGrowth:  1.5,
		SpiritGrowth:   1.5,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "flame-lash", MinLv: 1},
			{Key: "scorch", MinLv: 3},
			{Key: "molten-burst", MinLv: 4},
		},
	},

	// —————————————————————————————————————————————
	// ripple_chip
	// —————————————————————————————————————————————
	"ripple_chip": {
		Elements:       []Element{Water},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 1.5,
		DefenseGrowth:  1.5,
		SpiritGrowth:   3,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "squirt", MinLv: 1},
			{Key: "sweet-heal", MinLv: 3},
		},
	},

	// —————————————————————————————————————————————
	// fayluna
	// —————————————————————————————————————————————
	"fayluna": {
		Elements:       []Element{Astral},
		BaseHealth:     30,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 1.5,
		DefenseGrowth:  1.5,
		SpiritGrowth:   3,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bash", MinLv: 0},
			{Key: "celestial-beam", MinLv: 1},
			{Key: "astral-veil", MinLv: 3},
		},
	},

	// —————————————————————————————————————————————
	// stonebound_sentinel (mini boss)
	// —————————————————————————————————————————————
	"stonebound_sentinel": {
		Elements:       []Element{Earth},
		BaseHealth:     35,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    7,
		BaseSpirit:     4,
		BaseSpeed:      3,
		HPGrowth:       6,
		StrengthGrowth: 2.5,
		DefenseGrowth:  3.5,
		SpiritGrowth:   1,
		SpeedGrowth:    1,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "boulder-bash", MinLv: 0},
			{Key: "iron-bulwark", MinLv: 1},
			{Key: "earthen-grasp", MinLv: 2},
			{Key: "stone-spire", MinLv: 3},
		},
	},

	// —————————————————————————————————————————————
	// gravebound_husk (mini boss)
	// —————————————————————————————————————————————
	"gravebound_husk": {
		Elements:       []Element{Earth, Dark},
		BaseHealth:     25,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "decay-claw", MinLv: 0},
			{Key: "rot-tide", MinLv: 2},
		},
	},

	// —————————————————————————————————————————————
	// boneoak_zombie (mini boss)
	// —————————————————————————————————————————————
	"boneoak_zombie": {
		Elements:       []Element{Earth, Dark},
		BaseHealth:     25,
		BaseMana:       4,
		BaseStrength:   5,
		BaseDefense:    5,
		BaseSpirit:     5,
		BaseSpeed:      5,
		HPGrowth:       4,
		StrengthGrowth: 2,
		DefenseGrowth:  2,
		SpiritGrowth:   2,
		SpeedGrowth:    2,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "bone-spewer", MinLv: 0},
			{Key: "oak-shackle", MinLv: 2},
		},
	},

	// —————————————————————————————————————————————
	// mycera, the Hollowbinder (mini boss)
	// —————————————————————————————————————————————
	"mycera": {
		Elements:       []Element{Earth, Dark},
		BaseHealth:     35,
		BaseMana:       4,
		BaseStrength:   6,
		BaseDefense:    5,
		BaseSpirit:     4,
		BaseSpeed:      5,
		HPGrowth:       6,
		StrengthGrowth: 3,
		DefenseGrowth:  3,
		SpiritGrowth:   2,
		SpeedGrowth:    3,
		Evasion:        0,
		AbilityTemplates: []CharacterAbilityTemplate{
			{Key: "leaf-slash", MinLv: 0},
			{Key: "nature-blessing", MinLv: 1},
			{Key: "earthquake", MinLv: 2},
		},
	},
}
