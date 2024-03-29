package knabjam

type Card struct {
	id          uint
	name        string // Not sure this is necessary
	Label       string
	Number      uint
	Description string
	Rarity      string // This should be one of 4 presets: common|rare|legendary|singularity
	ImageFile   string // This should be a path to a real file. Probably a url to IPFS
	AssetId     string // The cardano asset id, ie d1b9ee6fc8879fa7262e73dc22690dc938c0b4cb1911507170330d0259756b69
	PolicyId    string // The cardano policy id, ie d1b9ee6fc8879fa7262e73dc22690dc938c0b4cb1911507170330d02
}
