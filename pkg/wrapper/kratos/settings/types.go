package settings

type UpdateProfileAPIBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Traits    Traits `json:"traits"`
}

type ChangePasswordAPIBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
}

type ToggleTOTPAPIBody struct {
	FlowID     string `json:"flowID"`
	CsrfToken  string `json:"csrf_token"`
	TOTPCode   string `json:"totp_code"`
	TOTPUnlink bool   `json:"totp_unlink"`
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
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	ImgURL       string `json:"img_url,omitempty"`
	PhoneNumber  string `json:"phone_number"`
	InviteStatus string `json:"invite_status"`
	Verified     bool   `json:"verified"`
	Role         string `json:"role"`
	Created_At   string `json:"created_at"`
	TOTP_Enabled bool   `json:"totp_enabled"`
}
