package client

import "encoding/json"

const (
	defaultHost    = "https://circleci.com"
	defaultPort    = 443
	defaultVersion = "v1"
)

// Client is for configuring settings to interact with the
// CircleCI API to make requests
type Client struct {
	Token   string `json:"token"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Version string `json:"version"`
}

// NewClient takes an auth token and returns a new Client object
func NewClient(token string) *Client {
	return &Client{
		Token:   token,
		Host:    defaultHost,
		Port:    defaultPort,
		Version: defaultVersion,
	}
}

func (client *Client) String() string {
	jsonClient, _ := json.Marshal(client)
	return string(jsonClient)
}
