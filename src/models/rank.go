package models

// Ranks is a collection of rank objects
type Ranks []Rank

// Rank is a rank object
type Rank struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Skills Skills `json:"skills,omitempty"`
}

// Save Rank
func (r *Rank) Save() error {
	//TODO: implement save functionality
	return nil
}

// Delete Rank
func (r *Rank) Delete() error {
	//TODO: implement delete functionality
	return nil
}
