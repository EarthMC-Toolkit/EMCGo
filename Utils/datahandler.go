package utils

import (
	//"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

var htmlErrorCodes = []int{
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409,
	410, 411, 412, 413, 414, 415, 416, 417, 418, 421,
	422, 423, 424, 425, 426, 428, 429, 431, 451, 500,
	501, 502, 503, 504, 505,
}

// func NewOAPI(mapName string) *OAPI {
// 	domain := fmt.Sprintf("https://api.earthmc.net/v2/%s", mapName)
// 	urls := map[string]string{
// 		"towns":     fmt.Sprintf("%s/towns", domain),
// 		"nations":   fmt.Sprintf("%s/nations", domain),
// 		"residents": fmt.Sprintf("%s/residents", domain),
// 	}

// 	return &OAPI{
// 		domain: domain,
// 		urls:   urls,
// 	}
// }

// func (o *OAPI) FetchSingle(endpointType, item string) interface{} {
// 	url := o.urls[endpointType] + "/" + item

// 	return reqJSON(url)
// }

// func (o *OAPI) FetchAll(endpointType string) interface{} {
// 	url := o.urls[endpointType]

// 	return reqJSON(url)
// }
