package circleci

import (
	"net/url"
)

// CircleIface - CircleCI Client interface
type CircleIface interface {
	Me() (*Me, *APIResponse)
	CheckoutKeys(username, project string) ([]*CheckoutKey, *APIResponse)
	GetCheckoutKey(username, project, fingerprint string) (*CheckoutKey, *APIResponse)
	DeleteCheckoutKey(username, project, fingerprint string) (*APIMessageResponse, *APIResponse)
	RecentBuilds(params url.Values) ([]*Build, *APIResponse)
	NewBuild(username, project string, body *BuildPostBody) (*Build, *APIResponse)
	BuildBranch(username, project, branch string, body *BuildPostBody) (*Build, *APIResponse)
	GetBuild(username, project string, buildNum int) (*Build, *APIResponse)
	RetryBuild(username, project string, buildNum int) (*Build, *APIResponse)
	CancelBuild(username, project string, buildNum int) (*Build, *APIResponse)
	BuildArtifacts(username, project string, buildNum int) ([]*Artifact, *APIResponse)
	BuildTests(username, project string, buildNum int) (*BuildTests, *APIResponse)
	String() string
	Projects() ([]*Project, *APIResponse)
	ProjectFollow(username, project string) (*ProjectFollow, *APIResponse)
	ProjectUnfollow(username, project string) (*ProjectFollow, *APIResponse)
	ProjectClearCache(username, project string) (*ProjectClearCache, *APIResponse)
	ProjectRecentBuilds(username, project string, params url.Values) ([]*Build, *APIResponse)
	ProjectRecentBuildsBranch(username, project, branch string, params url.Values) ([]*Build, *APIResponse)
	ProjectEnable(username, project string) (*Project, *APIResponse)
	ProjectSettings(username, project string) (*Project, *APIResponse)
	EnvVars(username, project string) ([]*EnvVar, *APIResponse)
	AddEnvVar(username, project, name, value string) (*EnvVar, *APIResponse)
	GetEnvVar(username, project, name string) (*EnvVar, *APIResponse)
	DeleteEnvVar(username, project, name string) (*APIMessageResponse, *APIResponse)
}
