// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package api

import (
	"net/http"
)

// HandleGetLoginFlow handles the GET request for the login flow.
func HandleGetLoginFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("not implemented"))
}

// HandlePostLoginFlow handles the POST request for the login flow.
func HandlePostLoginFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("not implemented"))
}
