package recovery

type SubmitRecoveryAPIBody struct {
	CsrfToken string `json:"csrf_token"`
	FlowID    string `json:"flowID"`
	Email     string `json:"email"`
}
