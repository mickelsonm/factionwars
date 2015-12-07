package models

// Players is a collection of player objects
type Players []Player

// Player is a player object
type Player struct {
	Member
	Faction   Faction `json:"faction"`
	Email     string  `json:"email"`
	Username  string  `json:"user_name"`
	Password  string  `json:"-"`
	IPAddress string  `json:"ip_address"`
}

// Save Player
func (p *Player) Save() error {
	//TODO: implement save functionality
	return nil
}

// Delete Player
func (p *Player) Delete() error {
	//TODO: implement delete functionality
	return nil
}
