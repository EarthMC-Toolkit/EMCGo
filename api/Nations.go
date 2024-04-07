package api

import (
	"errors"
	"github.com/earthmc-toolkit/emcgo/api/parser"
	"github.com/earthmc-toolkit/emcgo/structs"

	"github.com/samber/lo"
)

type NationList map[string]structs.Nation

func (nations NationList) Get(name string) (*structs.Nation, error) {
	n, found := nations[name]
	return &n, lo.Ternary(found, nil, errors.New("Could not find nation with name: "+name))
}

type Nations struct {
	MapName string
}

func (nations *Nations) All() (NationList, error) {
	var err error
	var mapRes structs.MapResponse

	mapRes, err = FetchData(nations.MapName)
	if err != nil {
		return nil, err
	}

	var towns map[string]structs.Town
	towns, err = parser.ParseTowns(mapRes.Sets.Towny)
	if err != nil {
		return nil, err
	}

	return parser.ParseNations(towns), nil
}

func (nations *Nations) Get(name string) (*structs.Nation, error) {
	all, err := nations.All()
	if err != nil {
		return nil, err
	}

	nation, found := all[name]
	return &nation, lo.Ternary(found, nil, errors.New("Could not find nation with name: "+name))
}
