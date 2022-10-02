package logout

type SubmitLogoutBody struct {
	// LogoutToken can be used to perform logout using AJAX.
	LogoutToken string `json:"logout_token"`
	// LogoutURL can be opened in a browser to sign the user out.  format: uri
	LogoutUrl string `json:"logout_url"`
}
