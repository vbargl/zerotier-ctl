package res

import (
	"context"
	"iter"
	"net/http"

	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

func ListMembers(ctx context.Context, client *zt.Client, netid string) (iter.Seq[*zt.ControllerNetworkMember], error) {
	httpRes, err := client.MemberListNetworkMembers2(ctx, zt.ZTNetworkID(netid))
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseMemberListNetworkMembers2Response(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		return nil, NotFoundError{}
	case http.StatusUnauthorized:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}

	return func(yield func(*zt.ControllerNetworkMember) bool) {
		for _, member := range res.JSON200.Data {
			if !yield(&member) {
				break
			}
		}
	}, nil
}

func ListMemberIds(ctx context.Context, client *zt.Client, netid string) (iter.Seq[string], error) {
	all, err := ListMembers(ctx, client, netid)
	if err != nil {
		return nil, err
	}

	return func(yield func(string) bool) {
		for member := range all {
			if !yield(member.Id) {
				break
			}
		}
	}, nil
}

func CreateMember(ctx context.Context, client *zt.Client, netid, memid string, body zt.ControllerNetworkMemberRequest) (*zt.ControllerNetworkMember, error) {
	httpRes, err := client.NetworkMemberPostNetworkMember(ctx, zt.ZTNetworkID(netid), zt.ZTAddress(memid), body)
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkMemberPostNetworkMemberResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}

func GetMember(ctx context.Context, client *zt.Client, netid, nodeid string) (*zt.ControllerNetworkMember, error) {
	httpRes, err := client.NetworkMemberGetNetworkMember(ctx, zt.ZTNetworkID(netid), zt.ZTAddress(nodeid))
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkMemberGetNetworkMemberResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}

func UpdateMember(ctx context.Context, client *zt.Client, netid, memid string, body zt.ControllerNetworkMemberRequest) (*zt.ControllerNetworkMember, error) {
	httpRes, err := client.NetworkMemberPostNetworkMember(ctx, zt.ZTNetworkID(netid), zt.ZTAddress(memid), body)
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkMemberPostNetworkMemberResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}

func DeleteMember(ctx context.Context, client *zt.Client, netid, nodeid string) (*zt.ControllerNetworkMember, error) {
	httpRes, err := client.NetworkMemberDelNetworkMember(ctx, zt.ZTNetworkID(netid), zt.ZTAddress(nodeid))
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkMemberDelNetworkMemberResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}
