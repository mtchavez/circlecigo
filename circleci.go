package circleci

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultHost    = "https://circleci.com"
	defaultPort    = 443
	defaultScheme  = "https"
	defaultVersion = "v1"
)

var (
	defaultURL    = &url.URL{Host: defaultHost, Scheme: defaultScheme, Path: "/api/" + defaultVersion}
	defaultLogger = log.New(os.Stderr, "[circleci]", log.LstdFlags)
)

// Client is for configuring settings to interact with the
// Client API to make requests
type Client struct {
	BaseURL *url.URL
	Logger  *log.Logger
	Token   string
	client  *http.Client
}

// NewClient takes an auth token and returns a new Client object
func NewClient(token string) *Client {
	return &Client{BaseURL: defaultURL, Token: token, Logger: defaultLogger}
}

func (client *Client) String() string {
	clientDetails := map[string]interface{}{
		"base_url": client.baseURL().String(),
		"token":    client.Token,
	}
	jsonClient, _ := json.Marshal(clientDetails)
	return string(jsonClient)
}

func (client *Client) baseURL() *url.URL {
	if client.BaseURL != nil {
		return client.BaseURL
	}
	return defaultURL
}
