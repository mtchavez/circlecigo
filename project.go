package circleci

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// BuildLimitMax is the max builds to return from the API
	BuildLimitMax = 100
)

var (
	// ValidBuildFilters is the list of valid build statuses you can filter by
	ValidBuildFilters = []string{"completed", "successful", "failed", "running"}
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

// ProjectRecentBuilds /project/:username/:project endpoint to get all the
// recent builds for the found project.
// Takes url.Values to set limit of returned builds (no more than BuildLimitMax)
// and filter to filter builds by status
// and an offset of number of builds to page through
func (client *Client) ProjectRecentBuilds(username, project string, params url.Values) ([]*Build, *APIResponse) {
	builds := []*Build{}
	if params == nil {
		params = url.Values{}
	}
	client.verifyBuildsParams(&params)
	path := fmt.Sprintf("/project/%s/%s", username, project)
	apiResp := client.request(http.MethodGet, path, params, nil, &builds)
	return builds, apiResp

}

// verifyBuildsParams ensures limit param is not greater than the max
// and that the filter is a valid option
func (client *Client) verifyBuildsParams(params *url.Values) {
	limitParam := params.Get("limit")
	var parsedLimit int
	var limitErr error
	if limitParam != "" {
		parsedLimit, limitErr = strconv.Atoi(limitParam)
	}
	if limitErr != nil || (parsedLimit > 100 || parsedLimit < 0) {
		client.Logger.Printf("Invalid limit, defaulting to %d", BuildLimitMax)
		params.Set("limit", strconv.Itoa(BuildLimitMax))
	}
	filterParam := params.Get("filter")
	if filterParam != "" && !validBuildFilter(filterParam) {
		client.Logger.Printf("Invalid filter %s, defaulting to empty", filterParam)
		params.Del("filter")
	}
}

// validBuildFilter verifies that a filter is in the ValidBuildFilters
func validBuildFilter(filter string) bool {
	valid := false
	for _, validFilter := range ValidBuildFilters {
		if validFilter == filter {
			valid = true
			break
		}
	}
	return valid
}
