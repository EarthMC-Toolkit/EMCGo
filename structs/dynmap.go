package structs

type EndpointData struct {
	Map			Endpoints 	`json:"map"`
	Players		Endpoints	`json:"players"`
	Config		Endpoints	`json:"config"`
}

type Endpoints struct {
	Nova		string 		`json:"nova"`
	Aurora		string		`json:"aurora"`
}

type MapResponse struct {
	Sets 		MapSets		`json:"sets"`
}

type MapSets struct {
	Markers   	Markerset	`json:"markers"`
	WorldBorder Markerset 	`json:"worldborder.markerset"`
	Towny		Markerset	`json:"townyPlugin.markerset"`
}

type Markerset struct {
    Areas 		any			`json:"areas"`
    Label 		string		`json:"label"`
    Markers 	any			`json:"markers"`
    Lines		any			`json:"lines"`
}