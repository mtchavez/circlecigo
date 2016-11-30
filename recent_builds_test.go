package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_RecentBuilds__Unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testMux.HandleFunc("/recent-builds", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	builds, apiResp := testClient.RecentBuilds(nil)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 0 {
		t.Errorf("Expected no builds but got %d", len(builds))
	}
}

func TestClient_RecentBuilds(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testMux.HandleFunc("/recent-builds", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"build_num": 12345, "branch": "master"}]`)
	})
	builds, apiResp := testClient.RecentBuilds(nil)
	if !apiResp.Success() {
		t.Errorf("Expected a successful response with no errors")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	fmt.Printf("%v", builds)
	if len(builds) == 0 {
		t.Errorf("Expected builds but got %d", len(builds))
	}
	build := builds[0]
	if build.BuildNum != 12345 {
		t.Errorf("Expected build number to be %d but got %d", 12345, build.BuildNum)
	}
	if build.Branch != "master" {
		t.Errorf("Expected build branch to be %s but got %d", "master", build.BuildNum)
	}
}
