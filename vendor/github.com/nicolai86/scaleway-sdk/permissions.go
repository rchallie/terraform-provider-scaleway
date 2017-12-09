package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Permissions represents the response of GET /permissions
type Permissions map[string]PermCategory

// PermCategory represents Permissions's fields
type PermCategory map[string][]string

// PermissionDefinition represents the permissions
type PermissionDefinition struct {
	Permissions Permissions `json:"permissions"`
}

// GetPermissions returns the permissions
func (s *API) GetPermissions() (*PermissionDefinition, error) {
	resp, err := s.GetResponsePaginate(AccountAPI, fmt.Sprintf("tokens/%s/permissions", s.Token), url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := s.handleHTTPError([]int{http.StatusOK}, resp)
	if err != nil {
		return nil, err
	}
	var permissions PermissionDefinition

	if err = json.Unmarshal(body, &permissions); err != nil {
		return nil, err
	}
	return &permissions, nil
}
