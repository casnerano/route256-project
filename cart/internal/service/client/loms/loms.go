package loms

type client struct {
	baseURL string
}

func NewClient(baseURL string) *client {
	return &client{baseURL: baseURL}
}
