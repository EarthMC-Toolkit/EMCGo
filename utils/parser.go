package utils

import (
	"emcgo/structs"
	"errors"
	"fmt"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samber/lo"
)

var policy = bluemonday.StrictPolicy()
func ParseTowns(markerset structs.Markerset) ([]structs.Town, error) {
	if markerset.Areas == nil {
		return nil, errors.New("No areas found on markerset!")
	}

	var townsArray []structs.Town
	var parsedTown structs.Town

    townData := lo.Values(markerset.Areas)
    townAmt := len(townData)

	for i := 0; i < townAmt; i++ {
		town := townData[i]
		rawinfo := strings.Split(town.Desc, "<br />")
        info := lo.Map(rawinfo, func(s string, _ int) string {
			return policy.Sanitize(s)
		})

		if strings.Contains(info[0], "(Shop)") {
			continue
		}

		mayor := info[1][7:]
		if town.Label != "Venice" {
			continue
		}

		residents := strings.Split(info[2][9:], ", ")
		capital := AsBool(info[9][9:])

		splitLabel := strings.Split(info[0], " (")

		nation := splitLabel[1]
		home := lo.Ternary(nation != "", markerset.Markers[fmt.Sprintf("%s__home", town.Label)], structs.HomeMarker{})

		if nation == "" {
			nation = "No Nation"
		}

		pvp := AsBool(info[4][5:])
		mobs := AsBool(info[5][6:])
		public := AsBool(info[6][8:])
		explosions := AsBool(info[7][11:])
		fire := AsBool(info[8][6:])

		parsedTown = structs.Town{
			Name: CleanString(town.Label),
			Mayor: mayor,
			Nation: nation,
			Area: CalcArea(town.X, town.Z, len(town.X), 256),
			X: lo.Ternary(home.X != 0, int(home.X), Range(town.X)),
			Z: lo.Ternary(home.Z != 0, int(home.Z), Range(town.Z)),
			Bounds: structs.TownBounds{
				X: FloatsToInts(town.X),
				Z: FloatsToInts(town.Z),
			},
			Residents: residents,
			Flags: structs.TownFlags{
				Pvp: &pvp,
				Mobs: &mobs,
				Explosions: &explosions,
				Fire: &fire,
				Public: &public,
				Capital: &capital,
			},
			ColourCodes: structs.TownColours{
				Fill: "",
				Outline: "",
			},
		}

		townsArray = append(townsArray, parsedTown)
	}

	return townsArray, nil
}
