// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetRecoveryFlow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "recovery flow",
	})
}

func HandlePostRecoveryFlow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "recovery flow",
	})
}
