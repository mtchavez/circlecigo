package circleci

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
	"time"
)

func TestClient_NewBuild_buildBody(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		checkBodyContains(t, r, `{"revision":"","tag":"v1.1.2","parallel":4,"build_parameters":{"INTEGRATION_TESTS":"true"}}`)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testTriggerBuild)
	})
	buildBody := &BuildPostBody{
		Tag:      "v1.1.2",
		Parallel: 4,
		BuildParameters: map[string]string{
			"INTEGRATION_TESTS": "true",
		},
	}
	build, apiResp := testClient.NewBuild(testUsername, testReponame, buildBody)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "not_running" {
		t.Errorf("Expected build Status to be %s but got %s", "not_running", build.Status)
	}
}

func TestClient_NewBuild(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testTriggerBuild)
	})
	build, apiResp := testClient.NewBuild(testUsername, testReponame, nil)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "not_running" {
		t.Errorf("Expected build Status to be %s but got %s", "not_running", build.Status)
	}
}

func TestClient_BuildBranch_buildBody(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/tree/%s", testUsername, testReponame, "master")
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		checkBodyContains(t, r, `{"revision":"","tag":"v1.1.2","parallel":4,"build_parameters":{"INTEGRATION_TESTS":"true"}}`)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testTriggerBuild)
	})
	buildBody := &BuildPostBody{
		Tag:      "v1.1.2",
		Parallel: 4,
		BuildParameters: map[string]string{
			"INTEGRATION_TESTS": "true",
		},
	}
	build, apiResp := testClient.BuildBranch(testUsername, testReponame, "master", buildBody)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "not_running" {
		t.Errorf("Expected build Status to be %s but got %s", "not_running", build.Status)
	}
}

func TestClient_BuildBranch(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/tree/%s", testUsername, testReponame, "master")
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testTriggerBuild)
	})
	build, apiResp := testClient.BuildBranch(testUsername, testReponame, "master", nil)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "not_running" {
		t.Errorf("Expected build Status to be %s but got %s", "not_running", build.Status)
	}
}

func TestClient_GetBuild_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
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
		checkMethod(t, r, http.MethodGet)
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
		checkMethod(t, r, http.MethodPost)
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
		checkMethod(t, r, http.MethodPost)
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

func TestClient_CancelBuild_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/cancel", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	build, apiResp := testClient.CancelBuild(testUsername, testReponame, testBuildNum)
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

func TestClient_CancelBuild(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/cancel", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "cancelled", "build_num": 1234}`)
	})
	build, apiResp := testClient.CancelBuild(testUsername, testReponame, testBuildNum)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if build.Status != "cancelled" {
		t.Errorf("Expected build Status to be %s but got %s", "cancelled", build.Status)
	}
	if build.BuildNum != 1234 {
		t.Errorf("Expected build BuildNum to be %d but got %d", 1234, build.BuildNum)
	}
}

func TestClient_BuildArtifacts_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/artifacts", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	artifacts, apiResp := testClient.BuildArtifacts(testUsername, testReponame, testBuildNum)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if len(artifacts) != 0 {
		t.Errorf("Expected no artifacts to exist")
	}
}

func TestClient_BuildArtifacts(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/artifacts", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"node_index": 0, "url": "https://circleci.com/my-path-to-artifact"}]`)
	})
	artifacts, apiResp := testClient.BuildArtifacts(testUsername, testReponame, testBuildNum)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(artifacts) == 0 {
		t.Errorf("Expected artifacts but didn't get any")
	}
	artifact := artifacts[0]
	if artifact.URL == "" {
		t.Errorf("Expected artifact URL to be set")
	}
}

func TestClient_BuildTests_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/tests", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	tests, apiResp := testClient.BuildTests(testUsername, testReponame, testBuildNum)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
	if len(tests.Tests) != 0 {
		t.Errorf("Expected no tests to exist")
	}
}

func TestClient_BuildTestsTests_notFound(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/tests", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message": "Build not found"}`)
	})
	_, apiResp := testClient.BuildTests(testUsername, testReponame, testBuildNum)

	if apiResp.Success() {
		t.Errorf("Expected response not to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status NotFound but got %v", apiResp.Response.StatusCode)
	}
	notFoundMessage := "Build not found"
	if apiResp.ErrorResponse.Message != notFoundMessage {
		t.Errorf("Expected not found message but got %s", apiResp.ErrorResponse.Message)
	}
}

func TestClient_BuildTests(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/%d/tests", testUsername, testReponame, testBuildNum)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
  "tests" : [ {
    "message" : "",
    "file" : "features/desktop/invitations.feature",
    "source" : "cucumber",
    "run_time" : 2.957513661,
    "result" : "success",
    "name" : "Accepting an invitation",
    "classname" : "Invitations"
  }, {
    "message" : null,
    "file" : "spec/lib/webfinger_spec.rb",
    "source" : "rspec",
    "run_time" : 0.011366,
    "result" : "success",
    "name" : "Webfinger#intialize sets account ",
    "classname" : "spec.lib.webfinger_spec"
  } ]
}`)
	})
	tests, apiResp := testClient.BuildTests(testUsername, testReponame, testBuildNum)

	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	allTests := tests.Tests
	if len(allTests) != 2 {
		t.Errorf("Expected 2 tests but got %d", len(allTests))
	}
	firstTest := allTests[0]
	if firstTest.File == "" {
		t.Errorf("Expected test File to be returned")
	}
	if firstTest.Source == "" {
		t.Errorf("Expected test Source to be returned")
	}
	if firstTest.Result == "" {
		t.Errorf("Expected test Result to be returned")
	}
}

func TestBuild_RunTime_buildTimeMillisEmptyWithStartTime(t *testing.T) {
	startTime := time.Now().Add(time.Minute * time.Duration(10))
	build := &Build{
		BuildTimeMillis: 0,
		Status:          "running",
		StartTime:       startTime.Format(time.RFC3339),
	}
	runTime := build.RunTime()
	runTimeString := fmt.Sprintf("%+v", runTime)
	matchTime := "9m59.*s"
	if matched, e := regexp.MatchString(matchTime, runTimeString); e != nil || !matched {
		t.Errorf("Expected to match %v but got %v", matchTime, runTimeString)
	}
}

func TestBuild_RunTime_finishedRunning(t *testing.T) {
	build := &Build{
		BuildTimeMillis: 1000 * 60 * 60,
		Status:          "failed",
		StopTime:        time.Now().Format(time.RFC3339),
	}
	runTime := build.RunTime()
	runTimeString := fmt.Sprintf("%+v", runTime)
	matchTime := "1h0m0s"
	if matched, e := regexp.MatchString(matchTime, runTimeString); e != nil || !matched {
		t.Errorf("Expected to match %v but got %v", matchTime, runTimeString)
	}
}

func TestBuild_RunTime_runningWithBadStartTime(t *testing.T) {
	build := &Build{
		BuildTimeMillis: 0,
		Status:          "running",
		StartTime:       "1/1/3030 19:56:42",
	}
	runTime := build.RunTime()
	runTimeString := fmt.Sprintf("%+v", runTime)
	matchTime := "0s"
	if matched, e := regexp.MatchString(matchTime, runTimeString); e != nil || !matched {
		t.Errorf("Expected to match %v but got %v", matchTime, runTimeString)
	}
}
