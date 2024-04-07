package structs

type Timestamps struct {
	Registered     float64  `json:"registered"`
	LastOnline     *float64 `json:"lastOnline,omitempty"`
	JoinedNationAt *float64 `json:"joinedNationAt,omitempty"`
}

type RawTownStatus struct {
	Public      bool `json:"isPublic"`
	Open        bool `json:"isOpen"`
	Neutral     bool `json:"isNeutral"`
	Capital     bool `json:"isCapital"`
	Overclaimed bool `json:"isOverClaimed"`
	Ruined      bool `json:"isRuined"`
}

type RawTownStats struct {
	NumTownBlocks uint32  `json:"numTownBlocks"`
	MaxTownBlocks uint32  `json:"maxTownBlocks"`
	NumResidents  uint16  `json:"numResidents"`
	Balance       float32 `json:"balance"`
}

type RawTownCoords struct {
	Bounds TownBounds `json:"townBlocks"`
	Home   []string   `json:"home"`
	Spawn  struct {
		World string `json:"world"`
		Pitch int32  `json:"pitch"`
		Yaw   int32  `json:"yaw"`
		Location
	} `json:"spawn"`
}

type RawTownRanks struct {
	Councillor *[]string `json:"Councillor,omitempty"`
	Builder    *[]string `json:"Builder,omitempty"`
	Recruiter  *[]string `json:"Recruiter,omitempty"`
	Police     *[]string `json:"Police,omitempty"`
	TaxExempt  *[]string `json:"Tax-exempt,omitempty"`
	Treasurer  *[]string `json:"Treasurer,omitempty"`
}

type RawFlagPerms struct {
	Pvp        bool `json:"pvp"`
	Explosions bool `json:"explosion"`
	Fire       bool `json:"fire"`
	Mobs       bool `json:"mobs"`
}

type RawRnaoPerms struct {
	Build   []bool `json:"buildPerms"`
	Destroy []bool `json:"destroyPerms"`
	Switch  []bool `json:"switchPerms"`
	ItemUse []bool `json:"itemUsePerms"`
}

type RawTownPerms struct {
	Flags RawFlagPerms `json:"flagPerms"`
	Rnao  RawRnaoPerms `json:"rnaoPerms"`
}

type RawTown struct {
	Name        string        `json:"name"`
	UUID        string        `json:"uuid"`
	Mayor       string        `json:"mayor"`
	Board       string        `json:"board"`
	Founder     string        `json:"founder"`
	HexColor    string        `json:"mapColorHexCode"`
	Nation      *string       `json:"nation,omitempty"`
	Residents   []string      `json:"residents"`
	Timestamps  Timestamps    `json:"timestamps"`
	Status      RawTownStatus `json:"status"`
	Stats       RawTownStats  `json:"stats"`
	Coordinates RawTownCoords `json:"coordinates"`
	Ranks       RawTownRanks  `json:"ranks"`
	Trusted     *[]string     `json:"trusted,omitempty"`
	Outlaws     *[]string     `json:"outlaws,omitempty"`
	Perms       RawTownPerms  `json:"perms"`
}
