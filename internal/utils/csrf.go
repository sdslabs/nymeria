// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sdslabs/nymeria/internal/config"
)

// GenerateStatelessCSRF returns a token valid for 2 minutes
func GenerateCSRFToken(userID string) (string, error) {
	csrfSecretKey := []byte(config.AppConfig.CSRFSecret)
	timestamp := time.Now().Unix()
	payload := fmt.Sprintf("%s:%d", userID, timestamp)

	mac := hmac.New(sha256.New, csrfSecretKey)
	mac.Write([]byte(payload))
	signature := mac.Sum(nil)

	rawToken := fmt.Sprintf("%s:%d:%s", userID, timestamp, base64.URLEncoding.EncodeToString(signature))
	return base64.URLEncoding.EncodeToString([]byte(rawToken)), nil
}

func ValidateCSRFToken(userID, token string) bool {
	csrfSecretKey := []byte(config.AppConfig.CSRFSecret)
	decoded, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return false
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 3 {
		return false
	}

	tokenUserID := parts[0]
	timestampStr := parts[1]
	sigBase64 := parts[2]

	if tokenUserID != userID {
		return false
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return false
	}

	maxAge := time.Duration(config.AppConfig.CSRFMaxAge) * time.Minute

	// Check expiration
	if time.Since(time.Unix(timestamp, 0)) > maxAge {
		return false
	}

	// Recompute HMAC
	payload := fmt.Sprintf("%s:%d", userID, timestamp)
	mac := hmac.New(sha256.New, csrfSecretKey)
	mac.Write([]byte(payload))
	expectedSig := base64.URLEncoding.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expectedSig), []byte(sigBase64))
}
