// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"context"
	"fmt"

	"github.com/dashotv/fae"
	"github.com/dashotv/runic/internal/parser"
)

type ParserService struct {
	client *Client
}

// NewParser makes a new client for accessing Parser services.
func NewParserService(client *Client) *ParserService {
	return &ParserService{
		client: client,
	}
}

type ParserParseRequest struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

type ParserParseResponse struct {
	*Response
	Result *parser.TorrentInfo `json:"result"`
}

func (s *ParserService) Parse(ctx context.Context, req *ParserParseRequest) (*ParserParseResponse, error) {
	result := &ParserParseResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetQueryParam("title", fmt.Sprintf("%v", req.Title)).
		SetQueryParam("type", fmt.Sprintf("%v", req.Type)).
		Get("/parser/parse")
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

type ParserTitleRequest struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

type ParserTitleResponse struct {
	*Response
	Result *parser.TorrentInfo `json:"result"`
}

func (s *ParserService) Title(ctx context.Context, req *ParserTitleRequest) (*ParserTitleResponse, error) {
	result := &ParserTitleResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetQueryParam("title", fmt.Sprintf("%v", req.Title)).
		SetQueryParam("type", fmt.Sprintf("%v", req.Type)).
		Get("/parser/title")
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

type ParserBatchRequest struct {
	Batch *Batch `json:"batch"`
}

type ParserBatchResponse struct {
	*Response
	Result []*BatchResult `json:"result"`
}

func (s *ParserService) Batch(ctx context.Context, req *ParserBatchRequest) (*ParserBatchResponse, error) {
	result := &ParserBatchResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		Post("/parser/batch")
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
