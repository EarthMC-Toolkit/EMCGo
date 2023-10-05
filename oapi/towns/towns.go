package towns

import (
	"emcgo/utils"
	"emcgo/structs"
	"fmt"
)

func Get(name string) (structs.OAPITown, error) {
	town, err := utils.JsonRequest[structs.RawTown](fmt.Sprintf("/towns/%s", name), false)

	if err != nil { 
		return structs.OAPITown{}, err
	}
	
	bal := int32(town.Stats.Balance)
	
	return structs.OAPITown{
		Name: town.Name,
		UUID: &town.UUID,
		Founder: &town.Founder,
		Mayor: town.Mayor,
		Board: &town.Board,
		Balance: &bal,
		Timestamps: &town.Timestamps,
		Trusted: town.Trusted,
		Outlaws: town.Outlaws,
		Bounds: town.Coordinates.Bounds,
	}, nil
}