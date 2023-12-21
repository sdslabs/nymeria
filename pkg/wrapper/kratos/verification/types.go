package verification

type SubmitVerificationBody struct {
	CsrfToken string `json:"csrf_token"`
	FlowID    string `json:"flowID"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}
