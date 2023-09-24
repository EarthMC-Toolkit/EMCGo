package structs

type Timestamps struct {
	Registered		float64
	LastOnline		*float64
	JoinedNationAt	*float64
}

type RawTownStatus struct {
	Public			bool 		`json:"isPublic"`
	Open			bool 		`json:"isOpen"`
	Neutral			bool 		`json:"isNeutral"`
	Capital			bool		`json:"isCapital"`
	Overclaimed		bool		`json:"isOverClaimed"`
	Ruined			bool 		`json:"isRuined"`
}

type RawTownStats struct {
	NumTownBlocks	int16 		`json:"numTownBlocks"`
	MaxTownBlocks	int16 		`json:"maxTownBlocks"`
	NumResidents	int16 		`json:"numResidents"`
	Balance			float32 	`json:"balance"`
}

type RawTownCoords struct {
	Home 			[]string	`json:"home"`
	Spawn 			any 		`json:"spawn"`
}

type RawTownRanks struct {
	Councillor		[]string 	`json:"Councillor,omitempty"`
	Builder			[]string 	`json:"Builder,omitempty"`
	Recruiter		[]string 	`json:"Recruiter,omitempty"`
	Police			[]string 	`json:"Police,omitempty"`
	TaxExempt		[]string 	`json:"Tax-exempt,omitempty"`
	Treasurer		[]string 	`json:"Treasurer,omitempty"`
}

// TODO: Implement this
type RawTownPerms struct {

}

type RawTown struct {
	Name 			string			`json:"name"`
	UUID 			string			`json:"uuid"`
	Mayor 			string			`json:"mayor"`
	Board 			string			`json:"board"`
	Founder 		string			`json:"founder"`
	HexColor 		string			`json:"mapColorHexCode"`
	Nation 			string			`json:"nation,omitempty"`
	Residents		[]string 		`json:"residents"`
	Timestamps		Timestamps 		`json:"timestamps"`
	Status			RawTownStatus	`json:"status"`
	Stats			RawTownStats	`json:"stats"`
	Coordinates		RawTownCoords	`json:"coordinates"`
	Ranks			RawTownRanks	`json:"ranks"`
	Trusted			[]string		`json:"trusted"`
	Perms			any				`json:"perms"`
}