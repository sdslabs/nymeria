package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/log"
	"github.com/sdslabs/nymeria/pkg/db"
)

func HandleAppAuthorization(c *gin.Context) {
	var body AccessProfileRequest
	err := c.BindJSON(&body)
	if err != nil {
		log.ErrorLogger("Unable to process json body", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unable to process json body",
		})
		c.Abort()
	}
	app, err := db.GetApplication(body.ClientKey, body.ClientSecret)
	if err != nil {
		log.ErrorLogger("Unable to get application", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Internal Server Error",
		})
		c.Abort()
	}
	if app.RedirectURL != body.RedirectURL {
		log.ErrorLogger("Redirect URL does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Redirect URL does not match",
		})
		c.Abort()
	}
	if app.ClientKey != body.ClientKey {
		log.ErrorLogger("Client Key does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unauthorized",
		})
		c.Abort()
	}
	if app.ClientSecret != body.ClientSecret {
		log.ErrorLogger("Client Secret does not match", err)
		errCode, _ := strconv.Atoi(strings.Split(err.Error(), " ")[0])
		c.JSON(errCode, gin.H{
			"error":   strings.Split(err.Error(), " ")[1],
			"message": "Unauthorized",
		})
		c.Abort()
	}
	c.Next()
}
