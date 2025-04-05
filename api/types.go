package api

import "time"

type ApplicationPostBody struct {
	Name           string `json:"name"`
	RedirectURL    string `json:"redirect_url"`
	AllowedDomains string `json:"allowed_domains"`
	Organization   string `json:"organization"`
}

type ApplicationPutBody struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RedirectURL    string `json:"redirect_url"`
	AllowedDomains string `json:"allowed_domains"`
	Organization   string `json:"organization"`
}

type ApplicationBody struct {
	ID int `json:"id"`
}

type IdentityBody struct {
	Identity string `json:"identity"`
}

type VerifiableIdentityAddress struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	Id         *string    `json:"id,omitempty"`
	Status     string     `json:"status"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Value      string     `json:"value"`
	Verified   bool       `json:"verified"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	Via        string     `json:"via"`
}

type SecureAccessProfileRequest struct {
	RedirectURL string `json:"redirect_url"`
	ClientKey   string `json:"client_key"`
	Timestamp   string `json:"timestamp"` // Unix timestamp to prevent replay attacks
	Signature   string `json:"signature"` // HMAC signature of "client_key:timestamp:redirect_url"
}
