package structs

type OAPITown struct {
	Name 			string
	UUID 			*string
	Founder 		*string
	Mayor 			string
	Nation 			string
	Board 			*string
	Area			int32
	Bounds			TownBounds
	X				int32
	Z				int32
	Residents		[]string
	Timestamps		*Timestamps
	Balance			*int32
	Trusted			*[]string
	Outlaws			*[]string
	Perms			*any
	Wiki			*string
	ColourCodes 	TownColours
}

func (t OAPITown) GetName() string {
    return t.Name
}

func (t OAPITown) GetMayor() string {
    return t.Mayor
}

func (t OAPITown) GetNation() string {
    return t.Nation
}

func (t OAPITown) GetUUID() *string {
    return t.UUID
}

func (t OAPITown) GetBoard() *string {
    return t.Board
}

func (t OAPITown) GetBalance() *int32 {
    return t.Balance
}

func (t OAPITown) GetFounder() *string {
	return t.Founder
}

type TownBounds struct {
	X	[]int
	Z	[]int
}

type TownColours struct {
	Fill		string
	Outline		string
}

type TownFlags struct {
	Pvp 			*bool
	Explosions 		*bool
	Fire 			*bool
	Mobs 			*bool
	Public 			*bool
	Capital			*bool
}

type Town struct {
	Name 			string
	Mayor 			string
	Nation 			string
	Area			int
	Bounds			TownBounds
	X				int
	Z				int
	Residents		[]string
	Flags			TownFlags
	ColourCodes 	TownColours
}

func (t *Town) OnlineResidents() {
	
}