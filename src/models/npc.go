package models

// NPCS is a group of NPCs
type NPCS []NPC

// NPC is a non-playable monster (AI essentially)
type NPC struct {
	Member
	Faction Faction `json:"faction"`
}

// Save NPC
func (n *NPC) Save() error {
	//TODO: implement save functionality
	return nil
}

// Delete NPC
func (n *NPC) Delete() error {
	//TODO: implement delete functionality
	return nil
}
