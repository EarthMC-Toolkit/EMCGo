package structs

type EndpointData struct {
	Map				Endpoints 		`json:"map"`
	Players			Endpoints		`json:"players"`
	Config			Endpoints		`json:"config"`
}

type Endpoints struct {
	Nova			string 			`json:"nova"`
	Aurora			string			`json:"aurora"`
}

type MapResponse struct {
	Sets 			MapSets			`json:"sets"`
	Timestamp		int64			`json:"timestamp"`
}

type MapSets struct {
	Markers   		Markerset		`json:"markers"`
	WorldBorder 	Markerset 		`json:"chunky.markerset"`
	Towny			Markerset		`json:"townyPlugin.markerset"`
}

type Markerset struct {
    Areas 			map[string]MapArea	`json:"areas"`
    Label 			string			    `json:"label"`
    Markers 		any				    `json:"markers"`
    Lines			any				    `json:"lines"`
}

type MapArea struct {
    Label 			string			`json:"label"`
    X 				[]int32			`json:"x"`
    Z 				[]int32			`json:"z"`
    FillColour		string			`json:"fillcolor"`
	FillOpacity		float32			`json:"fillopacity"`
	OutlineColour 	string			`json:"color"`
    OutlineOpacity 	float32			`json:"opacity"`
    Desc			string			`json:"desc"`
}

type PlayersResponse struct {
    HasStorm 		bool			`json:"hasStorm"`
    Thundering 		bool			`json:"isThundering"`
    Count			int16			`json:"currentcount"`
    Players 		[]RawPlayer		`json:"players"`
    Updates 		[]UpdatedTile	`json:"updates"`
    ConfigHash 		int64			`json:"confighash"`
    ServerTime 		int64			`json:"servertime"`
    Timestamp 		int64			`json:"timestamp"`
}

type RawPlayer struct {		
    Name 			string			`json:"account"`
	Nickname 		string			`json:"name"`
	World 			string			`json:"world"`
	Location
}

type UpdatedTile struct {
    Name 			string			`json:""`
    Timestamp 		int64			`json:""`
}