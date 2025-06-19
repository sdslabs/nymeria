// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/config"
	"github.com/sdslabs/nymeria/internal/database"
	"github.com/sdslabs/nymeria/internal/logger"
	"github.com/sdslabs/nymeria/internal/middlewares"
)

func Start() {
	// Initialize global configuration
	config.Init()

	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	// Use custom logging middleware
	r.Use(logger.Middleware(logger.Logger))

	// CORS configuration
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // TODO: change in prod
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-CSRF-Token", "X-User-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours in seconds
	})

	r.Use(corsConfig)

	r.GET("/", func(c *gin.Context) {
		logger.Info().Msg("welcome")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "welcome",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		logger.Info().Msg("ping")
		logger.Info().Msg(config.AppConfig.DBHost)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})

	r.GET("/register", HandleGetRegistrationFlow)
	r.POST("/register", middlewares.CSRFMiddleware(), HandlePostRegistrationFlow)

	r.GET("/login", HandleGetLoginFlow)
	r.POST("/login", middlewares.CSRFMiddleware(), HandlePostLoginFlow)

	r.GET("/applications", HandleGetApplicationFlow)
	r.POST("/applications", HandleFetchAllApplicationsFlow)
	r.POST("/applications/:id", HandleFetchApplicationByIDFlow)

	r.GET("/applications/create", HandleGetApplicationFlow)
	r.POST("/applications/create", middlewares.CSRFMiddleware(), HandleCreateApplicationFlow)

	r.GET("/applications/update", HandleGetApplicationFlow)
	r.POST("/applications/update", middlewares.CSRFMiddleware(), HandleUpdateApplicationFlow)

	r.GET("/applications/delete", HandleGetApplicationFlow)
	r.DELETE("/applications/delete/:id", middlewares.CSRFMiddleware(), HandleDeleteApplicationFlow)

	r.Run(":9898")
}
