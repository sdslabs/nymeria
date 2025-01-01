package helper

type KratosHttpResponseBody struct {
	ID             string              `json:"id"`
	Type           string              `json:"type"`
	ExpiresAt      string              `json:"expires_at"`
	IssuedAt       string              `json:"issued_at"`
	RequestURL     string              `json:"request_url"`
	UI             KratosHttpReponseUI `json:"ui"`
	OrganizationID *string             `json:"organization_id"`
	State          string              `json:"state"`
}

type KratosHttpReponseUI struct {
	Action   string                     `json:"action"`
	Method   string                     `json:"method"`
	Nodes    []KratosHttpResponseNode   `json:"nodes"`
	Messages []KratosHttpReponseMessage `json:"messages"`
}

type KratosHttpResponseNode struct {
	Type       string                       `json:"type"`
	Group      string                       `json:"group"`
	Attributes KratosHttpResponseAttributes `json:"attributes"`
	Messages   []KratosHttpReponseMessage   `json:"messages"`
	Meta       KratosHttpResponseMeta       `json:"meta"`
}

type KratosHttpResponseAttributes struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Value        interface{} `json:"value"` // Can be string, boolean, or other types
	Required     bool        `json:"required,omitempty"`
	Disabled     bool        `json:"disabled,omitempty"`
	NodeType     string      `json:"node_type"`
	Autocomplete string      `json:"autocomplete,omitempty"`
}

type KratosHttpReponseMessage struct {
	ID      int                    `json:"id"`
	Text    string                 `json:"text"`
	Type    string                 `json:"type"`
	Context map[string]interface{} `json:"context"`
}

type KratosHttpResponseMeta struct {
	Label *KratosHttpResponseLabel `json:"label"`
}

type KratosHttpResponseLabel struct {
	ID      int                    `json:"id"`
	Text    string                 `json:"text"`
	Type    string                 `json:"type"`
	Context map[string]interface{} `json:"context"`
}
