package admin

type Identity struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	Image_url    string `json:"img_url"`
	Active       bool   `json:"active"`
	Verified     bool   `json:"verified"`
	Role         string `json:"role"`
	Created_at   string `json:"created_at"`
	Totp_enabled bool   `json:"totp_enabled"`
}
