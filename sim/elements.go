package sim

// Element is one of your elemental types.
type Element string

const (
    Air      Element = "air"
    Astral   Element = "astral"
    Dark     Element = "dark"
    Earth    Element = "earth"
    Fire     Element = "fire"
    Ice      Element = "ice"
    Sol      Element = "sol"
    Water    Element = "water"
    Wild     Element = "wild"
    Electric Element = "electric"
)

// elementInfo holds strengths (i.e. resistances) and weaknesses.
type elementInfo struct {
    Strengths  []Element // elements this one resists/is strong against
    Weaknesses []Element // elements this one is weak to
}

// ElementChart mirrors your JS elementWeaknessMap.
var ElementChart = map[Element]elementInfo{
    Air: {
        Strengths:  []Element{Wild},
        Weaknesses: []Element{Earth, Electric},
    },
    Astral: {
        Strengths:  []Element{Fire, Earth, Water, Wild},
        Weaknesses: []Element{Dark},
    },
    Dark: {
        Strengths:  []Element{Astral, Electric},
        Weaknesses: []Element{Sol, Wild},
    },
    Earth: {
        Strengths:  []Element{Air, Electric},
        Weaknesses: []Element{Fire, Astral, Ice},
    },
    Fire: {
        Strengths:  []Element{Earth, Ice},
        Weaknesses: []Element{Water, Astral},
    },
    Ice: {
        Strengths:  []Element{Water, Earth},
        Weaknesses: []Element{Fire},
    },
    Sol: {
        Strengths:  []Element{Dark},
        Weaknesses: []Element{Water},
    },
    Water: {
        Strengths:  []Element{Fire, Sol},
        Weaknesses: []Element{Ice, Astral, Electric},
    },
    Wild: {
        Strengths:  []Element{Dark},
        Weaknesses: []Element{Astral, Air},
    },
    Electric: {
        Strengths:  []Element{Water, Air},
        Weaknesses: []Element{Earth, Dark},
    },
}

// GetWeaknesses returns the slice of Elements that `e` is weak to,
// or nil if `e` is unknown.
func GetWeaknesses(e Element) []Element {
    if info, ok := ElementChart[e]; ok {
        return info.Weaknesses
    }
    return nil
}

// GetResistances (or strengths) returns the slice of Elements that `e` resists.
func GetResistances(e Element) []Element {
    if info, ok := ElementChart[e]; ok {
        return info.Strengths
    }
    return nil
}