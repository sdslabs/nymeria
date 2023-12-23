package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	client "github.com/ory/client-go"
	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
)

func getResponse(method string, endpoint string, query *bytes.Buffer) (string, error) {
	req, _ := http.NewRequest(method, endpoint, query)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HandleRbac(c *gin.Context) {
	log.Logger.Debug("RBAC")
	cookie, err := c.Cookie("sdslabs_session")

	if err != nil {
		log.ErrorLogger("Initialize Rbac Failed", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Initialize Rbac failed.",
		})
		return
	}

	apiClient := client.NewAPIClient(config.KratosClientConfig)
	session, _, err := apiClient.V0alpha2Api.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		log.ErrorLogger("Invalid Cookie", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Initialize Rbac failed.",
		})
		return
	}

	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]

	queryRelationEndpoint := config.KetoReadURL + "/relation-tuples"
	query, _ := json.Marshal(map[string]interface{}{
		"namespace":  "accounts",
		"relation":   "view",
		"subject_id": role,
	})

	jsonQuery := bytes.NewBuffer(query)

	res, err := getResponse("GET", queryRelationEndpoint, jsonQuery)

	if err != nil {
		log.ErrorLogger("Failed to query keto", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Initialize Rbac failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "RBAC passed",
		"role":    role,
		"res":     res,
	})
}
