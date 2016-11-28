package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_Me_Unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testMux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	me, apiResp := testClient.Me()
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if me.Login != "" {
		t.Errorf("Expected no login returned but got %v", me.Login)
	}
}

func TestClient_Me(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testMux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"login": "mtchavez", "selected_email": "foo@bar.com"}`)
	})
	me, apiResp := testClient.Me()
	if !apiResp.Success() {
		t.Errorf("Expected a successful response with no errors")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK code but got %v", apiResp.Response.StatusCode)
	}
	if me.Login == "" {
		t.Errorf("Expected a login but nothing was set")
	}
}
