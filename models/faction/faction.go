package factions

// Factions ... is a collection of factions
type Factions []Faction

// Faction ... is a faction
type Faction struct {
	Name    string  `json:"name"`
	Leader  Member  `json:"leader,omitempty"`
	Members Members `json:"members,omitempty"`
}
