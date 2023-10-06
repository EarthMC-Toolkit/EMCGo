package utils

import (
	"emcgo/structs"
	"emcgo/api/parser"
	"emcgo/utils"
	"fmt"
	"testing"
)

const mapEndpoint = "/map/aurora/standalone/MySQL_markers.php?marker=_markers_/marker_earth.json"

func TestParsedTowns(t *testing.T) {
	mapRes, _ := utils.JsonRequest[structs.MapResponse](mapEndpoint, false)
	parsed, _ := parser.ParseTowns(mapRes.Sets.Towny)

	fmt.Println(utils.Prettify(parsed["fami"]))
}