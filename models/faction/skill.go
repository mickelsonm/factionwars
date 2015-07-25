package factions

// Skills ... is a collection of skills
type Skills []Skill

// Skill ... is a skill
type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}
