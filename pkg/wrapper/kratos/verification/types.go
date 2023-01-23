package verification

type SubmitVerificationBody struct {
	CsrfToken     string `json:"csrf_token"`
	FlowID        string `json:"flowID"`
	RecoveryToken string `json:"recovery_token"`
	Email         string `json:"email"`
	Method        string `json:"method"`
}