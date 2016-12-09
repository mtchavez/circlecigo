package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_Follow_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/follow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	follow, apiResp := testClient.Follow(testUsername, testReponame)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if follow.Followed {
		t.Errorf("Expected not to be following project")
	}
}

func TestClient_Follow(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/follow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"followed": true, "first_build": 1234}`)
	})
	follow, apiResp := testClient.Follow(testUsername, testReponame)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if !follow.Followed {
		t.Errorf("Expected to be following project")
	}
}

func TestClient_Unfollow_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/unfollow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	_, apiResp := testClient.Unfollow(testUsername, testReponame)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
}

func TestClient_Unfollow(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/unfollow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"followed": false}`)
	})
	follow, apiResp := testClient.Unfollow(testUsername, testReponame)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if follow.Followed {
		t.Errorf("Expected to not be following project")
	}
}

func TestClient_ClearCache_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/build-cache", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	_, apiResp := testClient.ClearCache(testUsername, testReponame)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
}

func TestClient_ClearCache(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/build-cache", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "build cache deleted"}`)
	})
	clearCache, apiResp := testClient.ClearCache(testUsername, testReponame)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if clearCache.Status != "build cache deleted" {
		t.Errorf("Expected status to be %s but got %s", "build cache deleted", clearCache.Status)
	}
}
