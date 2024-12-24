package helper

type HttpResponseBody struct {
	UI struct {
		Messages []struct {
			Text string `json:"text"`
		} `json:"messages"`
	} `json:"ui"`
}
