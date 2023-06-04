package settings

type SubmitSettingsAPIBody struct {
	Method     string `json:"method"`
	FlowID     string `json:"flowID"`
	CsrfToken  string `json:"csrf_token"`
	TOTPCode   string `json:"totp_code"`
	TOTPUnlink bool   `json:"totp_unlink"`
	Password   string `json:"password"`
	Traits     Traits `json:"Traits"`
}

type SubmitSettingsProfileRequest struct {
	Method    string                 `json:"method"`
	CsrfToken string                 `json:"csrf_token"`
	Traits    map[string]interface{} `json:"traits"`
}

type SubmitSettingsWithTOTPBody struct {
	Method     string `json:"method"`
	CsrfToken  string `json:"csrf_token"`
	TotpCode   string `json:"totp_code"`
	TotpUnlink bool   `json:"totp_unlink"`
}

type SubmitSettingsWithPasswordBody struct {
	Method    string `json:"method"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
}

type Traits struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	ImgURL       string `json:"img_url,omitempty"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Active       bool   `json:"active"`
	Created_At   string `json:"created_at"`
	TOTP_Enabled bool   `json:"totp_enabled"`
}
