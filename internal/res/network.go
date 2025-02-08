package res

import (
	"context"
	"iter"

	zt "github.com/vbargl/zerotier-ctl/internal/zerotier"
)

func ListNetworks(ctx context.Context, client *zt.Client) (iter.Seq[*zt.ControllerNetwork], error) {
	httpRes, err := client.NetworkReadNetworks2(ctx)
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkReadNetworks2Response(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		// continue
	case 401:
		return nil, UnauthorizedError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}

	return func(yield func(*zt.ControllerNetwork) bool) {
		for _, network := range res.JSON200.Data {
			if !yield(&network) {
				break
			}
		}
	}, nil
}

func ListNetworkIds(ctx context.Context, client *zt.Client) (iter.Seq[string], error) {
	all, err := ListNetworks(ctx, client)
	if err != nil {
		return nil, err
	}

	return func(yield func(string) bool) {
		for network := range all {
			if !yield(network.Id) {
				break
			}
		}
	}, nil
}

func CreateNetwork(ctx context.Context, client *zt.Client, body zt.ControllerNetworkRequest) (*zt.ControllerNetwork, error) {
	httpRes, err := client.RandomNetworkRandomNetwork(ctx, body)
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseRandomNetworkRandomNetworkResponse(httpRes)
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

func GetNetwork(ctx context.Context, client *zt.Client, netid string) (*zt.ControllerNetwork, error) {
	httpRes, err := client.NetworkReadNetwork(ctx, zt.ZTNetworkID(netid))
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkReadNetworkResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	case 404:
		return nil, NotFoundError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}

func UpdateNetwork(ctx context.Context, client *zt.Client, netid string, body zt.ControllerNetworkRequest) (*zt.ControllerNetwork, error) {
	httpRes, err := client.NetworkPostNetwork(ctx, zt.ZTNetworkID(netid), body)
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkPostNetworkResponse(httpRes)
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

func DeleteNetwork(ctx context.Context, client *zt.Client, netid string) (*zt.ControllerNetwork, error) {
	httpRes, err := client.NetworkDeleteNetwork(ctx, zt.ZTNetworkID(netid))
	if err != nil {
		return nil, HttpError{err}
	}

	res, err := zt.ParseNetworkDeleteNetworkResponse(httpRes)
	if err != nil {
		return nil, ParsingError{err}
	}

	switch res.StatusCode() {
	case 200:
		return res.JSON200, nil
	case 401:
		return nil, UnauthorizedError{}
	case 404:
		return nil, NotFoundError{}
	default:
		return nil, &UnexpectedResponseError{res.StatusCode(), res.Body}
	}
}
