package Map

import (
	"github.com/earthmc-toolkit/earthmc-go/utils"
	"github.com/earthmc-toolkit/earthmc-go/oapi"
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