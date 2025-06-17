// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetLoginFlow handles the GET request for the login flow.
func HandleGetLoginFlow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login flow",
	})
}

// HandlePostLoginFlow handles the POST request for the login flow.
func HandlePostLoginFlow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login flow",
	})
}
