package api

type ApplicationPostBody struct {
	Name           string `json:"name"`
	RedirectURL    string `json:"redirect_url"`
	AllowedDomains string `json:"allowed_domains"`
	Organisation   string `json:"organisation"`
}

type ApplicationPutBody struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RedirectURL    string `json:"redirect_url"`
	AllowedDomains string `json:"allowed_domains"`
	Organisation   string `json:"organisation"`
}

type ApplicationDeleteBody struct {
	ID int `json:"id"`
}
