package Map

import (
	"emcgo/api"
)

var Aurora = createMap("aurora")
var Nova = createMap("nova")

func createMap(name string) Map {
	return Map{
		Towns: api.Towns{
			MapName: name,
		},
		Nations: api.Nations{
			MapName: name,
		},
		Players: api.Players{},
	}
}

type Map struct {
	Name    		string
	Nations 		api.Nations
	Towns   		api.Towns
	Players			api.Players
}