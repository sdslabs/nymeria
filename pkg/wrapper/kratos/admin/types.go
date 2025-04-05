package admin

type Identity struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	ImgURL      string `json:"img_url,omitempty"`
	GithubID    string `json:"github_id,omitempty"`
	InvitedBy   string `json:"invited_by,omitempty"`
}
