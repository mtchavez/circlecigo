package circleci

import (
	"encoding/json"
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

// AddEnvVar calls the /project/:username/:project/envvar endpoint
// to set an ENV variable for the project builds
func (client *Client) AddEnvVar(username, project, name, value string) (*EnvVar, *APIResponse) {
	envVar := &EnvVar{}
	addVar := &EnvVar{Name: name, Value: value}
	body, _ := json.Marshal(addVar)
	path := fmt.Sprintf("/project/%s/%s/envvar", username, project)
	apiResp := client.request(http.MethodPost, path, nil, body, &envVar)
	return envVar, apiResp
}
