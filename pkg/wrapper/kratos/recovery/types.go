package recovery

type SubmitRecoveryAPIBody struct {
	FlowID     string `json:"flowID"`
	RecoveryToken	   string `json:"recovery_token"`
	CsrfToken		string `json:"csrf_token"`
	Email			string `json:"email"`
	Method			string `json:"method"`
	Code			string `json:"code"`
}