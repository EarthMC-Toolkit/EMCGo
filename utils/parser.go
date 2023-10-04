package utils

import (
	"emcgo/structs"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samber/lo"
)

var policy = bluemonday.StrictPolicy()
func ParseTowns(areas map[string]structs.MapArea) []structs.MapArea {
	var townsArray []structs.MapArea

    townData := lo.Values(areas)
    len := len(townData)

	for i := 0; i < len; i++ {
		town := townData[i]
		rawinfo := strings.Split(town.Desc, "<br />")
        info := lo.Map(rawinfo, func(s string, _ int) string {
			return policy.Sanitize(s)
		})

		if strings.Contains(info[0], "(Shop)") {
			continue
		}

		//fmt.Println(info)
		townsArray = append(townsArray, town)
	}

	return townsArray
}
