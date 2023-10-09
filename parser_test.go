package utils

import (
	"emcgo/api/parser"
	"emcgo/structs"
	"emcgo/utils"
	"fmt"
	"testing"
)

const mapEndpoint = "/map/aurora/standalone/MySQL_markers.php?marker=_markers_/marker_earth.json"

func TestParsedTowns(t *testing.T) {
	mapRes, _ := utils.JsonRequest[structs.MapResponse](mapEndpoint, false)

	towns, _ := parser.ParseTowns(mapRes.Sets.Towny)
	nations := parser.ParseNations(towns)

	fmt.Println(utils.Prettify(nations["Venice"]))
}