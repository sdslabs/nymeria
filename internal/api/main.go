// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sdslabs/nymeria/internal/database"
)

func Start() {
	if err := database.Init(); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// r.Use(log.LoggerMiddleware(log.Logger))
	r.Use(middleware.Logger)

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // TODO: change in prod
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours in seconds
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong! 🏓"))
	})

	r.Get("/register", HandleGetRegistrationFlow)
	r.Post("/register", HandlePostRegistrationFlow)

	r.Get("/login", HandleGetLoginFlow)
	r.Post("/login", HandlePostLoginFlow)

	http.ListenAndServe(":9898", r)
}
