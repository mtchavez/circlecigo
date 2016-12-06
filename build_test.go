package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_GetBuild_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	build, apiResp := testClient.GetBuild(testUsername, testReponame, testBuildNum)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "" {
		t.Errorf("Expected no status to be set")
	}
}

func TestClient_GetBuild(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "running", "steps": [{"name": "make test", "actions": [{"name": "Running make test"}]}]}`)
	})
	build, apiResp := testClient.GetBuild(testUsername, testReponame, testBuildNum)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "running" {
		t.Errorf("Expected build Status to be %s but got %s", "running", build.Status)
	}
	steps := build.Steps
	if len(steps) != 1 {
		t.Errorf("Expected build to have %d steps but got %d", 1, len(steps))
	}
	firstStep := steps[0]
	if firstStep.Name != "make test" {
		t.Errorf("Expected step name to be %s but got %s", "make test", firstStep.Name)
	}
	actions := firstStep.Actions
	if len(actions) != 1 {
		t.Errorf("Expected build step to have %d actions but got %d", 1, len(actions))
	}
	firstAction := actions[0]
	if firstAction.Name != "Running make test" {
		t.Errorf("Expected action name to be %s but got %s", "Running make test", firstAction.Name)
	}
}

func TestClient_RetryBuild_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/retry", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	build, apiResp := testClient.RetryBuild(testUsername, testReponame, testBuildNum)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "" {
		t.Errorf("Expected no status to be set")
	}
}

func TestClient_RetryBuild(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/retry", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "running", "build_num": 1234}`)
	})
	build, apiResp := testClient.RetryBuild(testUsername, testReponame, testBuildNum)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "running" {
		t.Errorf("Expected build Status to be %s but got %s", "running", build.Status)
	}
	if build.BuildNum != 1234 {
		t.Errorf("Expected build BuildNum to be %d but got %d", 1234, build.BuildNum)
	}
}
