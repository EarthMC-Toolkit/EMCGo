package structs

type Town struct {
	Name 			string
	UUID 			string
	Nation 			string
	Mayor 			string
	Board 			string
	Founder 		string
	Residents		[]string
	Timestamps		Timestamps
	//Status			RawTownStatus
	Balance			*int32
	Bounds			TownBounds
	X				int32
	Z				int32
	Trusted			[]string
	Outlaws			[]string
	Perms			any
	Wiki			*string
	ColourCodes 	struct {
		Fill		string
		Outline		string
	}
}

type TownBounds struct {
	X	[]int
	Z	[]int
}

func (t *Town) OnlineResidents() {
	
}