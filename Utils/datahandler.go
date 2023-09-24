package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

var htmlErrorCodes = []int{
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
	410, 411, 412, 413, 414, 415, 416, 417, 418, 421,
	422, 423, 424, 425, 426, 428, 429, 431, 451, 500,
	501, 502, 503, 504, 505,
}

type Endpoint struct {
	urls map[string]string
}

func NewEndpoint() *Endpoint {
	endpoints := "https://raw.githubusercontent.com/EarthMC-Toolkit/Toolkit-Website/main/endpoints.json"
	urls := reqJSON(endpoints).(map[string]interface{})
	urlMap := make(map[string]string)

	for key, value := range urls {
		urlMap[key] = value.(string)
	}

	return &Endpoint{
		urls: urlMap,
	}
}

func (e *Endpoint) Fetch(endpointType, mapName string) interface{} {
	url := e.urls[endpointType+mapName]

	return reqJSON(url)
}

func reqJSON(url string) interface{} {
	cached, found := c.Get(url)
	if found {
		return cached
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error fetching endpoint: %s\nResponse content is not valid JSON!", url)
	}
	defer resp.Body.Close()

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Sprintf("Error fetching endpoint: %s\nResponse content is not valid JSON!", url)
	}

	c.Set(url, result, cache.DefaultExpiration)
	return result
}

type OAPI struct {
	domain string
	urls   map[string]string
}

func NewOAPI(mapName string) *OAPI {
	domain := fmt.Sprintf("https://api.earthmc.net/v2/%s", mapName)
	urls := map[string]string{
		"towns":     fmt.Sprintf("%s/towns", domain),
		"nations":   fmt.Sprintf("%s/nations", domain),
		"residents": fmt.Sprintf("%s/residents", domain),
	}

	return &OAPI{
		domain: domain,
		urls:   urls,
	}
}

func (o *OAPI) FetchSingle(endpointType, item string) interface{} {
	url := o.urls[endpointType] + "/" + item

	return reqJSON(url)
}

func (o *OAPI) FetchAll(endpointType string) interface{} {
	url := o.urls[endpointType]

	return reqJSON(url)
}
