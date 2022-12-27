package oidc

type SubmitOIDCLoginAPIBody struct {
	FlowID     string `json:"flowID"`
	CsrfToken  string `json:"csrf_token"`
}
