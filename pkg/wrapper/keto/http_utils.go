package keto

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func MakeRequest(endpoint Endpoint, data map[string]interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest(endpoint.Method, endpoint.URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jsonBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var body map[string]interface{}
	json.Unmarshal(jsonBody, &body)

	return body, nil
}
