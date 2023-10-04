package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const REQ_DOMAIN  = "https://earthmc.net"
const REQ_TIMEOUT = 6 * time.Second

func SendRequest(endpoint string, skipCache bool) ([]byte, error) {
	if skipCache == true {
		randStr := RandomString(12)
		endpoint += randStr
	}

	url := fmt.Sprintf("%s%s", REQ_DOMAIN, endpoint)
	client := http.Client{ Timeout: REQ_TIMEOUT }

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func JsonRequest[T any](endpoint string, skipCache bool) (T, error) {
	var data T
	res, err := SendRequest(endpoint, skipCache)

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(res), &data)
	return data, nil
}