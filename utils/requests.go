package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const REQ_TIMEOUT = 6 * time.Second

func Request(url string) ([]byte, error) {
	client := http.Client{ Timeout: REQ_TIMEOUT }

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func JsonRequest[T interface{}](url string) (T, error) {
	return toJSON[T](Request(url))
}

func toJSON[T interface{}](res []byte, err error) (T, error) {
	var data T

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(res), &data)
	return data, nil
}

func OAPIRequest(endpoint string, skipCache bool) ([]byte, error) {
	if skipCache == true {
		randStr := RandomString(12)
		endpoint += randStr
	}

	url := fmt.Sprintf("%s%s", "https://api.earthmc.net/v2/aurora", endpoint)
	return Request(url)
}

func OAPIJsonRequest[T interface{}](endpoint string, skipCache bool) (T, error) {
	return toJSON[T](OAPIRequest(endpoint, skipCache))
}