package login

type Traits struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type SubmitLoginAPIBody struct {
	FlowID     string `json:"flowID"`
	CsrfToken  string `json:"csrf_token"`
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
}

type SubmitLoginWithMFABody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	TOTP      string `json:"totp"`
}
