package structs

type EndpointData struct {
	Map     Endpoints `json:"map"`
	Players Endpoints `json:"players"`
	Config  Endpoints `json:"config"`
}

type Endpoints struct {
	Nova   string `json:"nova"`
	Aurora string `json:"aurora"`
}

type MapResponse struct {
	Sets      MapSets `json:"sets"`
	Timestamp uint32  `json:"timestamp"`
}

type MapSets struct {
	Markers     Markerset `json:"markers"`
	WorldBorder Markerset `json:"chunky.markerset"`
	Towny       Markerset `json:"townyPlugin.markerset"`
}

type Markerset struct {
	Areas   map[string]MapArea    `json:"areas"`
	Markers map[string]HomeMarker `json:"markers"`
	Lines   any                   `json:"lines"`
	Label   string                `json:"label"`
}

type HomeMarker struct {
	Markup bool    `json:"markup"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Z      float64 `json:"z"`
	Icon   string  `json:"icon"`
	Dim    string  `json:"dim"`
	Label  string  `json:"label"`
	Desc   string  `json:"desc"`
}

type MapArea struct {
	Label          string    `json:"label"`
	X              []float64 `json:"x"`
	Z              []float64 `json:"z"`
	FillColour     string    `json:"fillcolor"`
	FillOpacity    float32   `json:"fillopacity"`
	OutlineColour  string    `json:"color"`
	OutlineOpacity float32   `json:"opacity"`
	Desc           string    `json:"desc"`
}

type PlayersResponse struct {
	HasStorm   bool          `json:"hasStorm"`
	Thundering bool          `json:"isThundering"`
	Count      uint16        `json:"currentcount"`
	Players    []RawPlayer   `json:"players"`
	Updates    []UpdatedTile `json:"updates"`
	ConfigHash int64         `json:"confighash"`
	ServerTime int64         `json:"servertime"`
	Timestamp  uint32        `json:"timestamp"`
}

type RawPlayer struct {
	Name     string `json:"account"`
	Nickname string `json:"name"`
	World    string `json:"world"`
	Location
}

type UpdatedTile struct {
	Name      string `json:"name"`
	Timestamp uint32 `json:"timestamp"`
}
