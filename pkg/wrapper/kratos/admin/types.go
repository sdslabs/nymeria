package admin

type Identity struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	ImgURL      string `json:"img_url,omitempty"`
}
