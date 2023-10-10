package api

import (
	"emcgo/structs"
	"emcgo/utils"
)

const DOMAIN = "https://earthmc.net/map/"
const MAP_DATA_ENDPOINT = "/standalone/MySQL_markers.php?marker=_markers_/marker_earth.json"

func BuildURL(mapName string) string {
	return DOMAIN + mapName + MAP_DATA_ENDPOINT
}

func FetchData(mapName string) (structs.MapResponse, error) {
	return utils.JsonRequest[structs.MapResponse](BuildURL(mapName))
}