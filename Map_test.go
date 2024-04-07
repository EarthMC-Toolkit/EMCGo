package Map

import (
	"fmt"
	"github.com/earthmc-toolkit/emcgo/oapi"
	"github.com/earthmc-toolkit/emcgo/utils"
	"testing"
)

func TestAllTowns(t *testing.T) {
	towns, err := Aurora.Towns.All()

	if err != nil {
		t.Fatalf(err.Error())
		return
	}

	fmt.Println(utils.Prettify(towns))
}

func TestNation(t *testing.T) {
	nation, _ := Aurora.Nations.Get("Venice")
	fmt.Println(utils.Prettify(nation))
}

func TestOAPI(t *testing.T) {
	town, _ := oapi.Town("Venice")
	fmt.Println(utils.Prettify(town))
}
