package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/wrapper/kratos/admin"
)

func HandleCreateIdentityFlow(c *gin.Context) {

	var t admin.Identity

	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process JSON body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process JSON body",
		})
		return
	}

	var mappedJsonIdentity map[string]interface{}

	data, err := json.Marshal(t)

	if err != nil {
		log.ErrorLogger("Unable to convert map to json", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to convert map to json",
		})
		return
	}

	err = json.Unmarshal(data, &mappedJsonIdentity)

	if err != nil {
		log.ErrorLogger("Unable to convert JSON to map", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to convert JSON to map",
		})
		return
	}

	createdIdentity, r, err := admin.CreateIdentityFlowWrapper(mappedJsonIdentity)

	if err != nil {
		log.ErrorLogger("Error while calling `AdminCreateIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"identity": createdIdentity.Id,
	})
}

func HandleGetIdentityFlow(c *gin.Context) {
	createdIdentity := c.Query("identity")
	getIdentity, r, err := admin.GetIdentityFlowWrapper(createdIdentity)

	if err != nil {
		log.ErrorLogger("Error while calling `AdminGetIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	jsonString, err := json.Marshal(getIdentity.Traits)

	if err != nil {
		log.ErrorLogger("Unable to convert map to json", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to convert map to json",
		})
		return
	}

	var identity admin.Identity

	err = json.Unmarshal(jsonString, &identity)

	if err != nil {
		log.ErrorLogger("Unable to convert JSON to map", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to convert JSON to map",
		})
		return
	}

	fmt.Fprintf(os.Stdout, "Identity details for id %v. Traits: %v\n", createdIdentity, identity)
	c.JSON(http.StatusOK, gin.H{
		"Identity": createdIdentity,
		"Traits":   identity,
	})
}

func HandleDeleteIdentityFlow(c *gin.Context) {

	var t IdentityBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process JSON body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process JSON body",
		})
		return
	}

	r, err := admin.DeleteIdentityFlowWrapper(t.Identity)

	if err != nil {
		log.ErrorLogger("Error while calling `AdminDeleteIdentity`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "INternal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "removed identity",
	})
}

func HandleListIdentity(c *gin.Context) {
	identities, r, err := admin.ListIdentityFlowWrapper()
	if err != nil {
		log.ErrorLogger("Error while calling `AdminListIdentities`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"identities": identities,
	})
}

func HandleBanIdentity(c *gin.Context) {
	var t IdentityBody
	err := c.BindJSON(&t)

	if err != nil {
		log.ErrorLogger("Unable to process JSON body", err)

		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   err.Error(),
			"message": "Unable to process JSON body",
		})
		return
	}

	id, r, err := admin.BanIdentityFlowWrapper(t.Identity)

	if err != nil {
		log.ErrorLogger("Error while calling `AdminPatchIdentities`", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"identities": id,
	})
}