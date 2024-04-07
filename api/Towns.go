package api

import (
	"errors"
	"github.com/earthmc-toolkit/emcgo/api/parser"
	"github.com/earthmc-toolkit/emcgo/structs"

	"github.com/samber/lo"
)

type TownList map[string]structs.Town

func (towns TownList) Get(name string) (*structs.Town, error) {
	t, found := towns[name]
	return &t, lo.Ternary(found, nil, errors.New("Could not find town with name: "+name))
}

type Towns struct {
	TownList
	MapName string
}

func (towns *Towns) All() (TownList, error) {
	mapRes, _ := FetchData(towns.MapName)
	return parser.ParseTowns(mapRes.Sets.Towny)
}

func (towns *Towns) Get(name string) (*structs.Town, error) {
	all, err := towns.All()
	if err != nil {
		return nil, err
	}

	town, found := all[name]
	return &town, lo.Ternary(found, nil, errors.New("Could not find town with name: "+name))
}
