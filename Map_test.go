package Map

import (
	"emcgo/utils"
	"emcgo/oapi"
	"fmt"
	"testing"
)

func TestNation(t *testing.T) {
	nation, _ := Aurora.Nations.Get("Venice")
	fmt.Println(utils.Prettify(nation))
}

func TestOAPI(t *testing.T) {
	town, _ := oapi.Town("Venice")
	fmt.Println(utils.Prettify(town))
}