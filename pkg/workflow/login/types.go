package login

type Traits struct {
	Email string `json:"email"`
	Name  struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
}

type SubmitLoginAPIBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
	Traits    Traits `json:"traits"`
}
