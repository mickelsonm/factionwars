package models

// Members is a collection of member objects
type Members []Member

// Member is a member object
type Member struct {
	Name   string `json:"name"`
	Rank   Rank   `json:"rank,omitempty"`
	Skills Skills `json:"skills,omitempty"`
}
