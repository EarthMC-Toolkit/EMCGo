package parser

import (
	"errors"
	"fmt"
	"github.com/earthmc-toolkit/emcgo/structs"
	"github.com/earthmc-toolkit/emcgo/utils"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/samber/lo"
)

var policy = bluemonday.NewPolicy()

func ParseTowns(markerset structs.Markerset) (map[string]structs.Town, error) {
	//#region Handle bad markerset + init vars
	if markerset.Areas == nil {
		return nil, errors.New("no areas found on markerset")
	}

	towns := make(map[string]structs.Town)
	areas := lo.Values(markerset.Areas)

	policy.AllowElements("a")
	policy.AllowAttrs("href").OnElements("a")
	//#endregion

	//#region Loop over areas and parse into Towns.
	townAmt := len(areas)
	for i := 0; i < townAmt; i++ {
		town := areas[i]
		parsed, skip := ParseTown(town, markerset.Markers)

		if !skip {
			towns[town.Label] = parsed
		}
	}
	//#endregion

	return towns, nil
}

func ParseTown(town structs.MapArea, markers map[string]structs.HomeMarker) (structs.Town, bool) {
	//#region Parse area description into vars.
	rawinfo := strings.Split(town.Desc, "<br />")
	info := lo.Map(rawinfo, func(s string, _ int) string {
		return policy.Sanitize(s)
	})

	firstInfoEl := info[0]
	if strings.Contains(firstInfoEl, "(Shop)") {
		return structs.Town{}, true
	}

	mayor := info[1][7:]
	if mayor == "" {
		return structs.Town{}, true
	}

	fullLabel := strings.Split(firstInfoEl, " (")
	fullLabelLen := len(fullLabel)

	var label string
	if fullLabelLen >= 3 {
		label = fullLabel[2]
	} else {
		label = fullLabel[1]
	}

	label = label[:len(label)-1]

	residents := strings.Split(info[2][9:], ", ")

	pvp := utils.AsBool(info[4][5:])
	mobs := utils.AsBool(info[5][6:])
	public := utils.AsBool(info[6][8:])
	explosions := utils.AsBool(info[7][11:])
	fire := utils.AsBool(info[8][6:])
	capital := utils.AsBool(info[9][9:])
	//#endregion

	//#region Parse nation name (and wiki if it exists).
	var nation = label
	var wiki string

	if strings.Contains(label, "href") {
		index := strings.Index(label, ">") + 1
		nation = label[index:]
		nation = strings.ReplaceAll(nation, "</a>", "")

		label = strings.ReplaceAll(label, "<a href=\"", "")

		if capital {
			wiki = label[:strings.Index(label, "\"")]
		}
	}
	//#endregion

	//#region Build and return new Town.
	home := lo.Ternary(nation != "", markers[fmt.Sprintf("%s__home", town.Label)], structs.HomeMarker{})
	return structs.Town{
		Name:   utils.CleanString(town.Label),
		Mayor:  mayor,
		Nation: utils.CleanString(strings.TrimSpace(nation)),
		Area:   utils.CalcArea(town.X, town.Z, uint16(len(town.X)), 256),
		X:      lo.Ternary(home.X != 0, int(home.X), utils.Range(town.X)),
		Z:      lo.Ternary(home.Z != 0, int(home.Z), utils.Range(town.Z)),
		Bounds: structs.TownBounds{
			X: utils.FloatsToInts(town.X),
			Z: utils.FloatsToInts(town.Z),
		},
		Residents: residents,
		Flags: structs.TownFlags{
			Pvp:        &pvp,
			Mobs:       &mobs,
			Explosions: &explosions,
			Fire:       &fire,
			Public:     &public,
			Capital:    &capital,
		},
		ColourCodes: structs.TownColours{
			Fill:    town.FillColour,
			Outline: town.OutlineColour,
		},
		Wiki: &wiki,
	}, false
	//#endregion
}

func ParseNations(towns map[string]structs.Town) map[string]structs.Nation {
	nations := make(map[string]structs.Nation)

	for _, town := range towns {
		name := town.Nation
		if name == "" {
			continue
		}

		//#region Create Nation if doesn't exist already.
		nation, exists := nations[name]
		if !exists {
			nation = structs.Nation{
				Name: name,
			}
		}
		//#endregion

		//#region Values to add/set regardless of existence.
		nation.Area += town.Area

		residents := append(nation.Residents, town.Residents...)
		nation.Residents = lo.Uniq(residents)

		if nation.Name == name {
			nation.Towns = append(nation.Towns, town.Name)
		}

		if *town.Flags.Capital {
			if town.Wiki != nil {
				nation.Wiki = town.Wiki
			}

			nation.Leader = town.Mayor
			nation.Capital = structs.NationCapital{
				Name: town.Name,
				X:    town.X,
				Z:    town.Z,
			}
		}
		//#endregion

		nations[name] = nation
	}

	return nations
}
