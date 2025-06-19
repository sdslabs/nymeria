package api

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateApplicationRequest struct {
	Name           string   `json:"name" binding:"required"`
	ApplicationURL string   `json:"application_url" binding:"required"`
	AllowedOrigins []string `json:"allowed_origins" binding:"required"`
	RedirectURIs   []string `json:"redirect_uris" binding:"required"`
}

type UpdateApplicationRequest struct {
	ApplicationID  string   `json:"application_id" binding:"required"`
	ApplicationURL string   `json:"application_url" binding:"omitempty"`
	AllowedOrigins []string `json:"allowed_origins" binding:"omitempty"`
	RedirectURIs   []string `json:"redirect_uris" binding:"omitempty"`
	NewKeyFlag     bool     `json:"new_key_flag" binding:"omitempty"`
}

type VerifyEmailRequest struct {
	Email       string `json:"email" binding:"required"`
	IsIITRCheck bool   `json:"is_iitr_check" binding:"omitempty"`
}

type VerifyEmailCodeRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
