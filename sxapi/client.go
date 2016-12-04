package sxapi

import (
	"json"
	"net/http"
)

type Client interface {
	RequestPublic(path string) (json.RawMessage, error)
	RequestPrivate(path string, userId int) (json.RawMessage, error)
}

// client sends HTTP requests to StackExchange API.
// It's safe for concurrent use by multiple goroutines.
type client struct {
	endpoint string
	tokens   TokenPool
}

func NewClient(endpoint string, tokens TokenPool) Client {
	return &client{endpoint: endpoint, tokens: tokens}
}

func (c *client) RequestPublic(path string) (json.RawMessage, error) {
	// TODO: Write stat
	token, err := c.tokens.GetAny()
	if err != nil {
		return nil, err
	}
	defer c.tokens.PutBack(token)

	resp := c.httpRequest(url)
}

func (c *client) RequestPrivate(path string, userId int) (json.RawMessage, error) {
	return nil, nil
}

func (c *client) httpRequest(url string) (*response, error) {
	httpResp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		return nil, err
	}

	resp := &response{}
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	// TODO: check resp.ErrorId & resp.ErrorMessage & resp.ErrorName

	return resp, nil
}

type response struct {
	ErrorId        int    `json:"error_id"`
	ErrorMessage   string `json:"error_message"`
	ErrorName      string `json:"error_name"`
	HasMore        bool   `json:"has_more"`
	QuotaMax       int    `json:"quota_max"`
	QuotaRemaining int    `json:"quota_remaining"`
	Items          json.RawMessage
}
