package zerotier

import (
	"context"
	"encoding/json"
	"net/http"
)

func WithAuthToken(token string) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, createAuthTokenRequestEditor(token))
		return nil
	}
}

func createAuthTokenRequestEditor(token string) RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		req.Header.Set(AuthTokenHeader, token)
		return nil
	}
}

func P[T any](v T) *T {
	return &v
}

func (t *ControllerNetworkMember_IpAssignments) String() string {
	body := ""
	_ = json.Unmarshal(t.union, &body)
	return body
}

func (t *ControllerNetwork_IpAssignmentPools_IpRangeEnd) String() string {
	body := ""
	_ = json.Unmarshal(t.union, &body)
	return body
}

func (t *ControllerNetwork_IpAssignmentPools_IpRangeStart) String() string {
	body := ""
	_ = json.Unmarshal(t.union, &body)
	return body
}

func (t *ControllerNetwork_IpAssignmentPools) String() string {
	start := t.IpRangeStart.String()
	end := t.IpRangeEnd.String()
	if start == "" && end == "" {
		return ""
	}
	return start + "-" + end
}

func (t *ControllerNetwork_Routes) String() string {
	target := t.Target.String()
	via := ""
	if t.Via.IsSpecified() {
		viaT := t.Via.MustGet()
		via = viaT.String()
	}

	if target == "" {
		return ""
	}

	return t.Target.String() + " via " + via
}

func (t *ControllerNetwork_Routes_Target) String() string {
	body := ""
	_ = json.Unmarshal(t.union, &body)
	return body
}

func (t *ControllerNetwork_Routes_Via) String() string {
	body := ""
	_ = json.Unmarshal(t.union, &body)
	return body
}
