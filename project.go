package circleci

import "fmt"

// ProjectFollow - CircleCI API project follow response
type ProjectFollow struct {
	Followed   bool `json:"followed"`
	FirstBuild int  `json:"first_build"`
}

// Follow calls the /project/:username/:project/follow endpoint to follow a project
func (client *Client) Follow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/follow", username, project)
	apiResp := client.request("GET", path, nil, nil, follow)
	return follow, apiResp
}

// Unfollow calls the /project/:username/:project/unfollow endpoint to unfollow a project
func (client *Client) Unfollow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/unfollow", username, project)
	apiResp := client.request("GET", path, nil, nil, follow)
	return follow, apiResp
}
