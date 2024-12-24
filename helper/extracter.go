package helper

import (
	"encoding/json"
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

func ExtractErrorMessage(r *http.Response) string {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.ErrorLogger("Error message extractor failed: ", err)
		return "Error"
	}

	var resp HttpResponseBody
	err = json.Unmarshal(body, &resp)

	if err != nil {
		log.ErrorLogger("Error message extractor failed: ", err)
		return "Error"
	}

	if len(resp.UI.Messages) == 0 {
		return "Error"
	}

	return resp.UI.Messages[0].Text
}
