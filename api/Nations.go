package api

import (
	"errors"
	"github.com/earthmc-toolkit/emcgo/api/parser"
	"github.com/earthmc-toolkit/emcgo/structs"

	"github.com/samber/lo"
)

type Nations struct {
	MapName string
}

func (nations *Nations) All() map[string]structs.Nation {
	mapRes, _ := FetchData(nations.MapName)
	towns, _ := parser.ParseTowns(mapRes.Sets.Towny)

	return parser.ParseNations(towns)
}

func (nations *Nations) Get(name string) (structs.Nation, error) {
	nation, exists := nations.All()[name]
	return nation, lo.Ternary(exists, nil, errors.New("Could not find nation with name: "+name))
}
