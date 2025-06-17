// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/database"
	"github.com/sdslabs/nymeria/internal/log"
)

func Start() {
	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	// Use custom logging middleware
	r.Use(log.LoggerMiddleware(log.Logger))

	// CORS configuration
	config := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // TODO: change in prod
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours in seconds
	})
	
	r.Use(config)

	r.GET("/", func(c *gin.Context) {
		log.Logger.Info().Msg("welcome")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "welcome",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		log.Logger.Info().Msg("ping")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", HandlePostRegistrationFlow)

	r.GET("/login", HandleGetLoginFlow)
	r.POST("/login", HandlePostLoginFlow)

	r.Run(":9898")
}
