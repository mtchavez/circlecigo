package circleci

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/joho/godotenv"
)

const TestToken = "abcd-1234-abcd-1234"

var (
	testMux      *http.ServeMux
	testServer   *httptest.Server
	testClient   *Client
	testUsername = "mtchavez"
	testReponame = "circlecigo"
	testBuildNum = 1
)

func init() {
	godotenv.Load()
}

func startTestServer() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	testClient = defaultClient
	defaultClient.BaseURL, _ = url.Parse(testServer.URL)
}

func stopTestServer() {
	if testServer != nil {
		testServer.Close()
	}
}

func checkMethod(t *testing.T, r *http.Request, expected string) {
	if r.Method != expected {
		t.Errorf("Expected %v method but got %v", expected, r.Method)
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient(TestToken)
	if client.Token != TestToken {
		t.Errorf("Expected Token to be set to %s but got %s", TestToken, client.Token)
	}
}
