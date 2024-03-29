package knabjam

type User struct {
	id        uint
	Name      string
	Address   string // Cardano wallet address like addr1qy2hmlj...
	Level     uint
	Cards     []Card
	Bio       string
	Pic       Card // Allow selecting from cards owned.
	Pro       bool // User subscribes/pays
	Following []User
	Role      string // TODO: Should be a slice of Roles: Collector|Creator
}
