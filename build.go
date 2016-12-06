package circleci

import (
	"fmt"
	"time"
)

// Build - CircleCI build response
type Build struct {
	VcsURL          string       `json:"vcs_url"`
	BuildURL        string       `json:"build_url"`
	BuildNum        int          `json:"build_num"`
	Branch          string       `json:"branch"`
	VcsRevison      string       `json:"vcs_revision"`
	CommitterName   string       `json:"committer_name"`
	CommitterEmail  string       `json:"committer_email"`
	Subject         string       `json:"subject"`
	Body            string       `json:"body"`
	Why             string       `json:"why"`
	DontBuild       string       `json:"dont_build"`
	QueuedAt        string       `json:"queued_at"`
	StartTime       string       `json:"start_time"`
	StopTime        string       `json:"stop_time"`
	BuildTimeMillis int          `json:"build_time_millis"`
	Username        string       `json:"username"`
	Reponame        string       `json:"reponame"`
	Lifecycle       string       `json:"lifecycle"`
	Outcome         string       `json:"outcome"`
	Status          string       `json:"status"`
	Steps           []*Step      `json:"steps"`
	RetryOf         int          `json:"retry_of"`
	PreviousBuild   *BuildStatus `json:"previous_build"`
}

// BuildStatus - CircleCI status for a build response
type BuildStatus struct {
	BuildNum int    `json:"build_num"`
	Status   string `json:"status"`
}

// Step is a step in the build
type Step struct {
	Name    string    `json:"name"`
	Actions []*Action `json:"actions"`
}

// Action is a single step action for a build
type Action struct {
	BashCommand        string     `json:"bash_command"`
	RunTimeMillis      int        `json:"run_time_millis"`
	Continue           string     `json:"continue"`
	Parallel           bool       `json:"parallel"`
	StartTime          *time.Time `json:"start_time"`
	Name               string     `json:"name"`
	Messages           []string   `json:"messages"`
	Step               int        `json:"step"`
	ExitCode           int        `json:"exit_code"`
	EndTime            *time.Time `json:"end_time"`
	Index              int        `json:"index"`
	Status             string     `json:"status"`
	Timedout           bool       `json:"timedout"`
	InfrastructureFail bool       `json:"infrastructure_fail"`
	Type               string     `json:"type"`
	Source             string     `json:"source"`
	Failed             bool       `json:"failed"`
}

// GetBuild calls the /project/:username/:reponame/:build_num endpoint to return a build
func (client *Client) GetBuild(username, project string, buildNum int) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/%d", username, project, buildNum)
	apiResp := client.request("GET", path, nil, nil, build)
	return build, apiResp
}

// RetryBuild calls the /project/:username/:reponame/:build_num/retry endpoint to retry a build
func (client *Client) RetryBuild(username, project string, buildNum int) (*Build, *APIResponse) {
	build := &Build{}
	path := fmt.Sprintf("project/%s/%s/%d/retry", username, project, buildNum)
	apiResp := client.request("GET", path, nil, nil, build)
	return build, apiResp
}
