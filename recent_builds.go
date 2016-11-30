package circleci

import "net/url"

func (client *Client) RecentBuilds(params url.Values) ([]*Build, *ApiResponse) {
	builds := []*Build{}
	apiResp := client.request("GET", "/recent-builds", params, nil, &builds)
	return builds, apiResp
}
