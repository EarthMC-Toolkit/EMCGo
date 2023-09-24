package towns

import (
	"emcgo/utils"
	"emcgo/structs"
	"fmt"
)

func Get(name string) (structs.RawTown, error) {
	town, err := utils.JsonRequest[structs.RawTown](fmt.Sprintf("/towns/%s", name), false)

	if err != nil { 
		return structs.RawTown{}, err
	}
	
	return town, nil
}