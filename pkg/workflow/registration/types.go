package registration

import "time"

type InitializeRegistration struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	ExpiresAt  time.Time `json:"expires_at"`
	IssuedAt   time.Time `json:"issued_at"`
	RequestURL string    `json:"request_url"`
	UI         UI        `json:"ui"`
}

type UI struct {
	Action string `json:"action"`
	Method string `json:"method"`
	Nodes  []Node `json:"nodes"`
}

type Node struct {
	Type       string        `json:"type"`
	Group      string        `json:"group"`
	Attributes Attributes    `json:"attributes"`
	Messages   []interface{} `json:"messages"`
	Meta       `json:"meta"`
}

type Attributes struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Required bool   `json:"required"`
	Disabled bool   `json:"disabled"`
	NodeType string `json:"node_type"`
}

type Meta struct {
	Label Label `json:"label"`
}

type Label struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type SubmitRegistrationBody struct {
	Method    string `json:"method"`
	Password  string `json:"password"`
	CsrfToken string `json:"csrf_token"`
	Data      Traits `json:"traits"`
}

type SubmitRegistrationAPIBody struct {
	FlowID    string `json:"flowID"`
	CsrfToken string `json:"csrf_token"`
	Password  string `json:"password"`
	Traits    Traits `json:"traits"`
}

type Traits struct {
	Email string `json:"email"`
	Name  struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
}
