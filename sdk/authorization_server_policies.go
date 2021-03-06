package sdk

import (
	"context"
	"fmt"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type (
	AuthorizationServerPolicy struct {
		Status      string            `json:"status,omitempty"`
		Priority    int               `json:"priority,omitempty"`
		Type        string            `json:"type,omitempty"`
		Description string            `json:"description,omitempty"`
		Name        string            `json:"name,omitempty"`
		Id          string            `json:"id,omitempty"`
		Conditions  *PolicyConditions `json:"conditions,omitempty"`
	}

	PolicyConditions struct {
		Clients *Whitelist `json:"clients,omitempty"`
	}

	Whitelist struct {
		Include []string `json:"include,omitempty"`
	}
)

func (m *ApiSupplement) DeleteAuthorizationServerPolicy(authServerId, id string) (*okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/policies/%s", authServerId, id)
	req, err := m.RequestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return m.RequestExecutor.Do(context.Background(), req, nil)
}

func (m *ApiSupplement) ListAuthorizationServerPolicies(authServerId string) ([]*AuthorizationServerPolicy, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/policies", authServerId)
	req, err := m.RequestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var auth []*AuthorizationServerPolicy
	resp, err := m.RequestExecutor.Do(context.Background(), req, &auth)
	return auth, resp, err
}

func (m *ApiSupplement) CreateAuthorizationServerPolicy(authServerId string, body AuthorizationServerPolicy, qp *query.Params) (*AuthorizationServerPolicy, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/policies", authServerId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.RequestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := body
	resp, err := m.RequestExecutor.Do(context.Background(), req, &authorizationServer)
	return &authorizationServer, resp, err
}

func (m *ApiSupplement) UpdateAuthorizationServerPolicy(authServerId, id string, body AuthorizationServerPolicy, qp *query.Params) (*AuthorizationServerPolicy, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/policies/%s", authServerId, id)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.RequestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := body
	resp, err := m.RequestExecutor.Do(context.Background(), req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}
	return &authorizationServer, resp, nil
}

func (m *ApiSupplement) GetAuthorizationServerPolicy(authServerId, id string, authorizationServerInstance AuthorizationServerPolicy) (*AuthorizationServerPolicy, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/policies/%s", authServerId, id)
	req, err := m.RequestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := authorizationServerInstance
	resp, err := m.RequestExecutor.Do(context.Background(), req, &authorizationServer)
	if err != nil {
		return nil, resp, err
	}
	return &authorizationServer, resp, nil
}
