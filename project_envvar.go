package circleci

import (
	"fmt"
	"net/http"
)

// EnvVar - CircleCI API ENV variable response
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// EnvVars calls the /project/:username/:project/envvar endpoint to get
// all ENV variables set for the project builds
func (client *Client) EnvVars(username, project string) ([]*EnvVar, *APIResponse) {
	envVars := []*EnvVar{}
	path := fmt.Sprintf("/project/%s/%s/envvar", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, &envVars)
	return envVars, apiResp
}
