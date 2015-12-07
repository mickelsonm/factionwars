package models

// Skills is a collection of skill objects
type Skills []Skill

// Skill is a skill object
type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Save Skill
func (s *Skill) Save() error {
	//TODO: implement save functionality
	return nil
}

// Delete Skill
func (s *Skill) Delete() error {
	//TODO: implement delete functionality
	return nil
}
