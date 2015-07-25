package factions

// Members ... a collection of members
type Members []Member

// Member ... is a member
type Member struct {
	Name   string `json:"name"`
	Rank   Rank   `json:"rank,omitempty"`
	Skills Skills `json:"skills,omitempty"`
}
