package models

// Factions is a collection of faction objects
type Factions []Faction

// Faction is a faction object
type Faction struct {
	Name    string  `json:"name"`
	Leader  Member  `json:"leader,omitempty"`
	Council Members `json:"council,omitempty"`
	Members Members `json:"members,omitempty"`
	Ranks   Ranks   `json:"ranks,omitempty"`
}

//Save Faction
func (f *Faction) Save() error {
	//TODO: implement save functionality
	return nil
}

//Delete Faction
func (f *Faction) Delete() error {
	//TODO: implement delete functionality
	return nil
}
