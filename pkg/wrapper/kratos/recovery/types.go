package recovery

type SubmitRecoveryAPIBody struct {
	CsrfToken string `json:"csrf_token"`
	FlowID    string `json:"flowID"`
	Code      string `json:"code"`
	Method    string `json:"method"`
}
