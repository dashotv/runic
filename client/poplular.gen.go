// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"context"
	"fmt"

	"github.com/dashotv/fae"
)

type PoplularService struct {
	client *Client
}

// NewPoplular makes a new client for accessing Poplular services.
func NewPoplularService(client *Client) *PoplularService {
	return &PoplularService{
		client: client,
	}
}

type PoplularIndexRequest struct {
	Interval string `json:"interval"`
}

type PoplularIndexResponse struct {
	*Response
	Result *PopularResponse `json:"result"`
	Total  int64            `json:"total"`
}

func (s *PoplularService) Index(ctx context.Context, req *PoplularIndexRequest) (*PoplularIndexResponse, error) {
	result := &PoplularIndexResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetPathParam("interval", fmt.Sprintf("%v", req.Interval)).
		Get("/poplular/{interval}")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}