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

type ApplicationBody struct {
	ID int `json:"id"`
}

type AccessProfileRequest struct {
	RedirectURL    string `json:"redirect_url"`
	ClientKey	  string `json:"client_key"`
	ClientSecret   string `json:"client_secret"`
}
