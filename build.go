package circleci

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
	RetryOf         int          `json:"retry_of"`
	PreviousBuild   *BuildStatus `json:"previous_build"`
}

// BuildStatus - CircleCI status for a build response
type BuildStatus struct {
	BuildNum int    `json:"build_num"`
	Status   string `json:"status"`
}
