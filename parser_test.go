package utils

import (
	"emcgo/structs"
	"emcgo/utils"
	"fmt"
	"testing"
)

const mapEndpoint = "/map/aurora/standalone/MySQL_markers.php?marker=_markers_/marker_earth.json"

func TestParsedTowns(t *testing.T) {
	mapRes, _ := utils.JsonRequest[structs.MapResponse](mapEndpoint, false)
	parsed := utils.ParseTowns(mapRes.Sets.Towny.Areas)

	fmt.Println(utils.Prettify(parsed[0]))
}