package registration

type SubmitRegistrationBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
	Traits    Traits `json:"traits"`
}

type Traits struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
