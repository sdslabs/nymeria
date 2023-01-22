package middleware

type AccessProfileRequest struct {
	RedirectURL  string `json:"redirect_url"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
}
