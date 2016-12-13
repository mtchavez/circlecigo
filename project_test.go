package circleci

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_ProjectFollow_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/follow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	follow, apiResp := testClient.ProjectFollow(testUsername, testReponame)
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

func TestClient_ProjectFollow(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/follow", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"followed": true, "first_build": 1234}`)
	})
	follow, apiResp := testClient.ProjectFollow(testUsername, testReponame)
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

func TestClient_ProjectRecentBuilds(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		checkQueryParam(t, r, "limit", "")
		checkQueryParam(t, r, "filter", "")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectRecentBuilds)
	})
	builds, apiResp := testClient.ProjectRecentBuilds(testUsername, testReponame, nil)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 1 {
		t.Errorf("Expected expected 1 build but got %v", len(builds))
	}
	build := builds[0]
	if build.Status != "failed" {
		t.Errorf("Expected status to be %s but got %s", "failed", build.Status)
	}
}

func TestClient_ProjectRecentBuilds_exceedLimitParam(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testParams := url.Values{}
	testParams.Set("limit", "300")
	path := fmt.Sprintf("/project/%s/%s", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		checkQueryParam(t, r, "limit", "100")
		checkQueryParam(t, r, "filter", "")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectRecentBuilds)
	})
	builds, apiResp := testClient.ProjectRecentBuilds(testUsername, testReponame, testParams)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 1 {
		t.Errorf("Expected expected 1 build but got %v", len(builds))
	}
	build := builds[0]
	if build.Status != "failed" {
		t.Errorf("Expected status to be %s but got %s", "failed", build.Status)
	}
}

func TestClient_ProjectRecentBuilds_invalidFilterParam(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testParams := url.Values{}
	testParams.Set("filter", "no-existent-status")
	path := fmt.Sprintf("/project/%s/%s", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		checkQueryParam(t, r, "limit", "")
		checkQueryParam(t, r, "filter", "")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectRecentBuilds)
	})
	builds, apiResp := testClient.ProjectRecentBuilds(testUsername, testReponame, testParams)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 1 {
		t.Errorf("Expected expected 1 build but got %v", len(builds))
	}
	build := builds[0]
	if build.Status != "failed" {
		t.Errorf("Expected status to be %s but got %s", "failed", build.Status)
	}
}

func TestClient_ProjectRecentBuildsBranch_notFound(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testBranch := "master"
	path := fmt.Sprintf("/project/%s/%s/tree/%s", testUsername, testReponame, testBranch)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		checkQueryParam(t, r, "limit", "")
		checkQueryParam(t, r, "filter", "")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	})
	builds, apiResp := testClient.ProjectRecentBuildsBranch(testUsername, testReponame, testBranch, nil)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 0 {
		t.Errorf("Expected expected no builds but got %v", len(builds))
	}
}

func TestClient_ProjectRecentBuildsBranch(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testBranch := "master"
	path := fmt.Sprintf("/project/%s/%s/tree/%s", testUsername, testReponame, testBranch)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		checkQueryParam(t, r, "limit", "")
		checkQueryParam(t, r, "filter", "")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectRecentBuilds)
	})
	builds, apiResp := testClient.ProjectRecentBuildsBranch(testUsername, testReponame, testBranch, nil)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(builds) != 1 {
		t.Errorf("Expected expected 1 build but got %v", len(builds))
	}
	build := builds[0]
	if build.Status != "failed" {
		t.Errorf("Expected status to be %s but got %s", "failed", build.Status)
	}
}

func TestClient_ProjectEnable(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/enable", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectResponse)
	})
	project, apiResp := testClient.ProjectEnable(testUsername, testReponame)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if project.Language != "Ruby" {
		t.Errorf("Expected Language to be %s but got %s", "Ruby", project.Language)
	}
	if project.Reponame != "circleci" {
		t.Errorf("Expected Reponame to be %s but got %s", "circleci", project.Reponame)
	}
	sshKeys := project.SSHKeys
	if len(sshKeys) != 4 {
		t.Errorf("Expected SSH keys to be %d but got %d", 4, len(sshKeys))
	}
	if !project.FeatureFlags["oss"] {
		t.Errorf("Expected FeatureFlags to be oss but got false")
	}
	if _, found := project.Branches["master"]; !found {
		t.Errorf("Expected Branches to have master branch")
	}
	if project.AWS.Keypair != nil {
		t.Errorf("Expected AWS.Keypair to be nil but got %+v", project.AWS.Keypair)
	}
}

func TestClient_ProjectSettings(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/settings", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testProjectResponse)
	})
	project, apiResp := testClient.ProjectSettings(testUsername, testReponame)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if project.Language != "Ruby" {
		t.Errorf("Expected Language to be %s but got %s", "Ruby", project.Language)
	}
	if project.Reponame != "circleci" {
		t.Errorf("Expected Reponame to be %s but got %s", "circleci", project.Reponame)
	}
	sshKeys := project.SSHKeys
	if len(sshKeys) != 4 {
		t.Errorf("Expected SSH keys to be %d but got %d", 4, len(sshKeys))
	}
	if !project.FeatureFlags["oss"] {
		t.Errorf("Expected FeatureFlags to be oss but got false")
	}
	if _, found := project.Branches["master"]; !found {
		t.Errorf("Expected Branches to have master branch")
	}
	if project.AWS.Keypair != nil {
		t.Errorf("Expected AWS.Keypair to be nil but got %+v", project.AWS.Keypair)
	}
}

func TestClient_Projects(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testMux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fmt.Sprintf("[%s]", testProjectResponse))
	})
	projects, apiResp := testClient.Projects()

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(projects) != 1 {
		t.Errorf("Expected 1 project but got %v", len(projects))
	}
	project := projects[0]
	if project.Language != "Ruby" {
		t.Errorf("Expected Language to be %s but got %s", "Ruby", project.Language)
	}
	if project.Reponame != "circleci" {
		t.Errorf("Expected Reponame to be %s but got %s", "circleci", project.Reponame)
	}
	sshKeys := project.SSHKeys
	if len(sshKeys) != 4 {
		t.Errorf("Expected SSH keys to be %d but got %d", 4, len(sshKeys))
	}
	if !project.FeatureFlags["oss"] {
		t.Errorf("Expected FeatureFlags to be oss but got false")
	}
	if _, found := project.Branches["master"]; !found {
		t.Errorf("Expected Branches to have master branch")
	}
	if project.AWS.Keypair != nil {
		t.Errorf("Expected AWS.Keypair to be nil but got %+v", project.AWS.Keypair)
	}
}
