package nations

// import (
// 	"emcgo/utils"
// 	"emcgo/structs"
// 	"fmt"
// )

// func Get(name string) (structs.NationInfo, error) {
// 	nation, err := utils.JsonRequest[structs.NationInfo](fmt.Sprintf("/nations/%s?", name), false)

// 	if err != nil { 
// 		return structs.NationInfo{}, err
// 	}
	
// 	return nation, nil
// }