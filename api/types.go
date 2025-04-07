package api

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
