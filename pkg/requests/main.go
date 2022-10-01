package requests

import (
	"encoding/json"
	"net/http"
)

func GetInitializeFlowJSON(url string, target interface{}) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}
	l
	defer resp.Body.Close()

	setCookie := resp.Header.Get("Set-Cookie")
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return "", err
	}

	return setCookie, nil
}
