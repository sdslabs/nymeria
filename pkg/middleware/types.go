package middleware

type SecureAccessProfileRequest struct {
	RedirectURL string `json:"redirect_url"`
	ClientKey   string `json:"client_key"`
	Timestamp   string `json:"timestamp"` // Unix timestamp to prevent replay attacks
	Signature   string `json:"signature"` // HMAC signature of "client_key:timestamp:redirect_url"
}
