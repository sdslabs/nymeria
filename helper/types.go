package helper

type KratosHttpResponseBody struct {
	ID               string               `json:"id"`
	Type             string               `json:"type"`
	ExpiresAt        string               `json:"expires_at"`
	IssuedAt         string               `json:"issued_at"`
	RequestURL       string               `json:"request_url"`
	UI               KratosHttpResponseUI `json:"ui"`
	OrganizationID   *string              `json:"organization_id"`
	State            string               `json:"state"`
	Active           string               `json:"active,omitempty"`
	TransientPayload interface{}          `json:"transient_payload,omitempty"`
}

type KratosHttpResponseUI struct {
	Action   string                      `json:"action"`
	Method   string                      `json:"method"`
	Nodes    []KratosHttpResponseNode    `json:"nodes"`
	Messages []KratosHttpResponseMessage `json:"messages"`
}

type KratosHttpResponseNode struct {
	Type       string                       `json:"type"`
	Group      string                       `json:"group"`
	Attributes KratosHttpResponseAttributes `json:"attributes"`
	Messages   []KratosHttpResponseMessage  `json:"messages"`
	Meta       KratosHttpResponseMeta       `json:"meta"`
}

type KratosHttpResponseAttributes struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Value        interface{} `json:"value"`
	Required     bool        `json:"required,omitempty"`
	Disabled     bool        `json:"disabled,omitempty"`
	NodeType     string      `json:"node_type"`
	Autocomplete string      `json:"autocomplete,omitempty"`
	Maxlength    int         `json:"maxlength,omitempty"`
	Pattern      string      `json:"pattern,omitempty"`
}

type KratosHttpResponseMessage struct {
	ID      int                    `json:"id"`
	Text    string                 `json:"text"`
	Type    string                 `json:"type"`
	Context map[string]interface{} `json:"context,omitempty"`
}

type KratosHttpResponseMeta struct {
	Label *KratosHttpResponseLabel `json:"label,omitempty"`
}

type KratosHttpResponseLabel struct {
	ID      int                    `json:"id"`
	Text    string                 `json:"text"`
	Type    string                 `json:"type"`
	Context map[string]interface{} `json:"context,omitempty"`
}
