package registration

type SubmitRegistrationBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
	Traits    Traits `json:"traits"`
}

type Traits struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	ImgURL      string `json:"img_url,omitempty"`
	GithubID    string `json:"github_id,omitempty"`
	InvitedBy   string `json:"invited_by,omitempty"`
}
