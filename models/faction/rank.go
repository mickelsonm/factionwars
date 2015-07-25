package factions

// Ranks ... is a collection of ranks
type Ranks []Rank

// Rank ... is a rank
type Rank struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Skills Skills `json:"skills,omitempty"`
}
