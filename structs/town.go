package structs

import "github.com/earthmc-toolkit/emcgo/utils"

type OAPITown struct {
	Name        string
	UUID        *string
	Founder     *string
	Mayor       string
	Nation      string
	Board       *string
	Area        uint
	Bounds      TownBounds
	X           int
	Z           int
	Residents   []string
	Timestamps  *Timestamps
	Balance     *uint
	Trusted     *[]string
	Outlaws     *[]string
	Perms       *any
	Wiki        *string
	ColourCodes TownColours
}

type TownBounds struct {
	X []int
	Z []int
}

type TownColours struct {
	Fill    string
	Outline string
}

func (colours TownColours) AsInts() (int, int) {
	fill := utils.HexToInt(colours.Fill)
	outline := utils.HexToInt(colours.Outline)

	return fill, outline
}

type TownFlags struct {
	Pvp        *bool
	Explosions *bool
	Fire       *bool
	Mobs       *bool
	Public     *bool
	Capital    *bool
}

type Town struct {
	Name        string
	Mayor       string
	Nation      string
	Area        uint
	Bounds      TownBounds
	X           int
	Z           int
	Residents   []string
	Flags       TownFlags
	ColourCodes TownColours
	Wiki        *string
}

func (t *Town) OnlineResidents() {

}
