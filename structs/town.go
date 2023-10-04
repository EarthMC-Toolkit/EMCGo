package structs

type TownBounds struct {
	X	[]int
	Z	[]int
}

type TownColours struct {
	Fill		string
	Outline		string
}

type Town struct {
	Name 			string
	UUID 			string
	Nation 			*string
	Mayor 			string
	Board 			string
	Founder 		string
	Residents		[]string
	Timestamps		Timestamps
	Balance			*int32
	Bounds			TownBounds
	X				int32
	Z				int32
	Trusted			[]string
	Outlaws			[]string
	Perms			any
	Wiki			*string
	ColourCodes 	TownColours
}

func (t Town) GetName() string {
    return t.Name
}

func (t Town) GetUUID() string {
    return t.UUID
}

func (t Town) GetFounder() string {
	return t.Founder
}

func (t Town) GetMayor() string {
    return t.Mayor
}

func (t Town) GetBoard() string {
    return t.Board
}

func (t Town) GetNation() *string {
    return t.Nation
}

func (t Town) GetBalance() *int32 {
    return t.Balance
}

func (t *Town) OnlineResidents() {
	
}