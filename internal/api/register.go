// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/database"
	"github.com/sdslabs/nymeria/internal/log"
)

// HandleGetRegistrationFlow handles the GET request for the registration flow.
func HandleGetRegistrationFlow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Registration flow",
	})
}

// HandlePostRegistrationFlow handles the POST request for the registration flow.
func HandlePostRegistrationFlow(c *gin.Context) {
	var req RegistrationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Logger.Err(err).Msg("invalid json")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Username is required",
		})
		return
	}
	if req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Password is required",
		})
		return
	}
	if req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email is required",
		})
		return
	}
	if req.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Phone number is required",
		})
		return
	}

	user := database.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	if result := database.DB.Create(&user); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "GitHub ID already exists",
			})
		} else {
			log.Logger.Err(result.Error).Msg("failed to insert user")
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
	})
}
