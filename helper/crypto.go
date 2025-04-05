package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Maximum allowed age for a signed request in seconds
const MaxSignatureAge = 300 // 5 minutes

// ValidateSignature verifies a HMAC signature and checks that it's not too old
func ValidateSignature(clientKey string, secretKey string, redirectURL string, signature string, timestamp int64) bool {
	// Check if timestamp is too old to prevent replay attacks
	currentTime := time.Now().Unix()
	if currentTime-timestamp > MaxSignatureAge {
		return false
	}

	// Calculate expected signature
	payload := clientKey + ":" + strconv.FormatInt(timestamp, 10) + ":" + redirectURL
	expectedSig := GenerateSignature(payload, secretKey)

	// Compare signatures (constant time comparison to prevent timing attacks)
	return hmac.Equal([]byte(signature), []byte(expectedSig))
}

// GenerateSignature creates an HMAC signature for the given payload using the secret key
func GenerateSignature(payload string, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
