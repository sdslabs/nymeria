package db

type Application struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RedirectURL    string `json:"redirect_url"`
	AllowedDomains string `json:"allowed_domains"`
	UpdatedAt      string `json:"updated_at"`
	Organization   string `json:"organization"`
	CreatedAt      string `json:"created_at"`
	ClientKey      string `json:"client_key"`
	ClientSecret   string `json:"client_secret"`
}
