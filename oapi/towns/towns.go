package towns

import (
	"emcgo/utils"
	"emcgo/oapi/structs"
	"fmt"
)

func Get(name string) (structs.Town, error) {
	town, err := utils.JsonRequest[structs.RawTown](fmt.Sprintf("/towns/%s", name), false)

	if err != nil { 
		return structs.Town{}, err
	}
	
	bal := int32(town.Stats.Balance)
	
	return structs.Town{
		Name: town.Name,
		UUID: town.UUID,
		Founder: town.Founder,
		Mayor: town.Mayor,
		Board: town.Board,
		Balance: &bal,
	}, nil
}