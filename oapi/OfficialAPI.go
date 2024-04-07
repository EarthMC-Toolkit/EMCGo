package oapi

import (
	"fmt"
	"github.com/earthmc-toolkit/emcgo/structs"
	"github.com/earthmc-toolkit/emcgo/utils"
)

func Town(name string) (structs.OAPITown, error) {
	town, err := utils.OAPIJsonRequest[structs.RawTown](fmt.Sprintf("/towns/%s", name), false)

	if err != nil {
		return structs.OAPITown{}, err
	}

	bal := uint(town.Stats.Balance)

	return structs.OAPITown{
		Name:       town.Name,
		UUID:       &town.UUID,
		Founder:    &town.Founder,
		Mayor:      town.Mayor,
		Board:      &town.Board,
		Balance:    &bal,
		Timestamps: &town.Timestamps,
		Trusted:    town.Trusted,
		Outlaws:    town.Outlaws,
		Bounds:     town.Coordinates.Bounds,
	}, nil
}
