package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sdslabs/nymeria/log"
)

func ExtractErrorCode(Error error) int {
	errCode, err := strconv.Atoi(strings.Split(Error.Error(), " ")[0])

	if err != nil {
		log.ErrorLogger("Error code extractor failed: ", err)
		return http.StatusInternalServerError
	}

	if errCode == 0 {
		return http.StatusInternalServerError
	}
	return errCode
}

func ExtractSuccessMessage(r *http.Response) string {

	if r == nil {
		log.ErrorLogger("Error message extractor failed: ", errors.New("response is nil"))
		return "Kratos Connection Refused"
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.ErrorLogger("Success message extractor failed: ", err)
		return "Error: " + err.Error()
	}

	var resp KratosHttpResponseBody
	err = json.Unmarshal(body, &resp)

	if err != nil {
		log.ErrorLogger("Info message extractor failed: ", err)
		return "Info message unmarshal error"
	}

	if len(resp.UI.Nodes) == 0 {
		log.Logger.Warn(string(body))
		log.ErrorLogger("Invalid response from Kratos.", err)
		return "Invalid response from Kratos."
	}

	msg := ""
	for _, node := range resp.UI.Nodes {
		for _, message := range node.Messages {
			if message.Type == "success" {
				msg += message.Text + ", "
			}
		}
	}
	for _, message := range resp.UI.Messages {
		if message.Type == "success" {
			msg += message.Text + ", "
		}
	}

	return strings.Trim(msg, ", ")
}

func ExtractInfoMessage(r *http.Response) string {

	if r == nil {
		log.ErrorLogger("Error message extractor failed: ", errors.New("response is nil"))
		return "Kratos Connection Refused"
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.ErrorLogger("Info message extractor failed: ", err)
		return "Error: " + err.Error()
	}

	var resp KratosHttpResponseBody
	err = json.Unmarshal(body, &resp)

	if err != nil {
		log.ErrorLogger("Info message extractor failed: ", err)
		return "Info message unmarshal error"
	}

	if len(resp.UI.Nodes) == 0 {
		log.Logger.Warn(string(body))
		log.ErrorLogger("Invalid response from Kratos.", err)
		return "Invalid response from Kratos."
	}

	msg := ""
	for _, node := range resp.UI.Nodes {
		for _, message := range node.Messages {
			if message.Type == "info" {
				msg += message.Text + ", "
			}
		}
	}
	for _, message := range resp.UI.Messages {
		if message.Type == "info" {
			msg += message.Text + ", "
		}
	}

	return strings.Trim(msg, ", ")
}

func ExtractErrorMessage(r *http.Response) string {

	if r == nil {
		log.ErrorLogger("Error message extractor failed: ", errors.New("response is nil"))
		return "Kratos Connection Refused"
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.ErrorLogger("Error message extractor failed: ", err)
		return "Error"
	}

	var resp KratosHttpResponseBody
	err = json.Unmarshal(body, &resp)

	if err != nil {
		log.ErrorLogger("Error message extractor failed: ", err)
		return "Error message unmarshal error"
	}

	if len(resp.UI.Nodes) == 0 {
		log.Logger.Warn(string(body))
		log.ErrorLogger("Invalid response from Kratos.", err)
		return "Invalid response from Kratos."
	}

	msg := ""
	for _, node := range resp.UI.Nodes {
		for _, message := range node.Messages {
			if message.Type == "error" {
				msg += message.Text + ", "
			}
		}
	}

	for _, message := range resp.UI.Messages {
		if message.Type == "error" {
			msg += message.Text + ", "
		}
	}

	return strings.Trim(msg, ", ")
}
