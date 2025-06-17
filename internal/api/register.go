// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sdslabs/nymeria/internal/database"
)

// HandleGetRegistrationFlow handles the GET request for the registration flow.
func HandleGetRegistrationFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("not implemented"))
}

// HandlePostRegistrationFlow handles the POST request for the registration flow.
func HandlePostRegistrationFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if req.Phone == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
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
			http.Error(w, "GitHub ID already exists", http.StatusConflict)
		} else {
			log.Fatalf("failed to insert user: %v", result.Error)
		}
	}

	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": "User registered successfully",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
