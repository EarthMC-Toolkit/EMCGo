package utils

import (
	"emcgo/structs"
	"fmt"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samber/lo"
)

var policy = bluemonday.StrictPolicy()
func ParseTowns(areas map[string]structs.MapArea) []structs.Town {
	var townsArray []structs.Town
	var parsedTown structs.Town

    townData := lo.Values(areas)
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

		// let split: string | string[] = info[0].split(" (")
		// split = (split[2] ?? split[1]).slice(0, -1)

		residents := strings.Split(info[2][9:], ", ")
		capital := AsBool(info[9][9:])

		parsedTown = structs.Town{
			Name: CleanString(town.Label),
			Mayor: mayor,
			Area: CalcArea(town.X, town.Z, len(town.X), 256),
			Residents: residents,
			Flags: structs.TownFlags{
				Capital: &capital,
			},
		}

		townsArray = append(townsArray, parsedTown)
	}

	return townsArray
}
