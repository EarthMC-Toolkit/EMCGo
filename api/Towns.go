package api

import (
	"errors"
	"github.com/earthmc-toolkit/emcgo/api/parser"
	"github.com/earthmc-toolkit/emcgo/structs"

	"github.com/samber/lo"
)

type Towns struct {
	MapName string
}

func (towns *Towns) All() (map[string]structs.Town, error) {
	mapRes, _ := FetchData(towns.MapName)
	return parser.ParseTowns(mapRes.Sets.Towny)
}

func (towns *Towns) Get(name string) (structs.Town, error) {
	all, err := towns.All()
	return all[name], lo.Ternary(err == nil, nil, errors.New("Could not find town with name: "+name))
}
