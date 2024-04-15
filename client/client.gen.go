// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"github.com/go-resty/resty/v2"
)

// Client is used to access Pace services.
type Client struct {
	// RemoteHost is the URL of the remote server that this Client should
	// access.
	RemoteHost string
	// Debug enables debug on Resty client
	Debug bool
	// Resty
	Resty *resty.Client

	// Services corresponding to the different endpoints (groups/routes)
	Indexers *IndexersService
	Parser   *ParserService
	Popular  *PopularService
	Releases *ReleasesService
	Sources  *SourcesService
}

// Set the debug flag
func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
	c.Resty.SetDebug(debug)
}

// New makes a new Client.
func New(remoteHost string) *Client {
	c := &Client{
		RemoteHost: remoteHost,
		Resty:      resty.New().SetBaseURL(remoteHost),
	}
	c.Indexers = NewIndexersService(c)
	c.Parser = NewParserService(c)
	c.Popular = NewPopularService(c)
	c.Releases = NewReleasesService(c)
	c.Sources = NewSourcesService(c)

	return c
}

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type Setting struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

type SettingsBatch struct {
	IDs   []string `json:"ids"`
	Name  string   `json:"name"`
	Value bool     `json:"value"`
}
