package circleci

import "fmt"

type ProjectFollow struct {
	Followed   bool `json:"followed"`
	FirstBuild int  `json:"first_build"`
}

func (client *Client) Follow(username, project string) (*ProjectFollow, *ApiResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/follow", username, project)
	apiResp := client.request("GET", path, nil, nil, follow)
	return follow, apiResp
}

func (client *Client) Unfollow(username, project string) (*ProjectFollow, *ApiResponse) {
	follow := &ProjectFollow{}
	path := fmt.Sprintf("/project/%s/%s/unfollow", username, project)
	apiResp := client.request("GET", path, nil, nil, follow)
	return follow, apiResp
}
