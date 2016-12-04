package circleci

import "net/url"

// RecentBuilds calls the /recent-builds endpoint to get recent builds across
// all followed projects. Can supply limit and offset params.
func (client *Client) RecentBuilds(params url.Values) ([]*Build, *APIResponse) {
	builds := []*Build{}
	apiResp := client.request("GET", "/recent-builds", params, nil, &builds)
	return builds, apiResp
}
