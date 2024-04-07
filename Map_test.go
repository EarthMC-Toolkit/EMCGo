package Map

import (
	"fmt"
	"github.com/earthmc-toolkit/emcgo/oapi"
	"github.com/earthmc-toolkit/emcgo/utils"
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
