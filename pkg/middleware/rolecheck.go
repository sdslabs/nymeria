package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyAdmin(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "admin" || role == "superadmin" {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": "Route requires admin role",
	})
	c.Abort()
}

func OnlyUser(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "user" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": "Route requires user role",
	})
	c.Abort()
}

func OnlySuperAdmin(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	identity := session.GetIdentity()
	traits := identity.GetTraits()
	role := traits.(map[string]interface{})["role"]
	if role == "superadmin" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": "Route requires superadmin role",
	})
	c.Abort()
}
