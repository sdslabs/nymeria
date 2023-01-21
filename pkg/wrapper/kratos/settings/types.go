package settings

type SubmitSettingsWithPasswordBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
}
type SubmitSettingsAPIBody struct {
	FlowID     string `json:"flowID"`
	Method     string `json:"method"`
	CsrfToken  string `json:"csrf_token"`
	TOTPcode   string `json:"totp_code"`
	TOTPUnlink bool   `json:"totp_unlink"`
}
