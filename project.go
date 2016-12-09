package circleci

import (
	"fmt"
	"net/http"
)

// ProjectFollow - CircleCI API project follow response
type ProjectFollow struct {
	Followed   bool `json:"followed"`
	FirstBuild int  `json:"first_build"`
}

// ProjectClearCache - CircleCI API response when clearing a project cache
type ProjectClearCache struct {
	Status string `json:"status"`
}

// Follow calls the /project/:username/:project/follow endpoint to follow a project
func (client *Client) Follow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/follow", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, follow)
	return follow, apiResp
}

// Unfollow calls the /project/:username/:project/unfollow endpoint to unfollow a project
func (client *Client) Unfollow(username, project string) (*ProjectFollow, *APIResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/unfollow", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, follow)
	return follow, apiResp
}

// ClearCache calls the /project/:username/:project/unfollow endpoint to unfollow a project
func (client *Client) ClearCache(username, project string) (*ProjectClearCache, *APIResponse) {
	clearCache := &ProjectClearCache{}
	path := fmt.Sprintf("/project/%s/%s/build-cache", username, project)
	apiResp := client.request(http.MethodDelete, path, nil, nil, clearCache)
	return clearCache, apiResp
}
