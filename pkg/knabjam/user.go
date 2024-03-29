package knabjam

type User struct {
	Name      string
	Address   string // Cardano wallet address like addr1qy2hmlj...
	Bio       string
	Role      string // TODO: Should be a slice of Roles: Collector|Creator
	Cards     []Card
	Following []User
	Pic       Card // Allow selecting from cards owned.
	id        uint
	Level     uint
	Pro       bool // User subscribes/pays
}
