package gonut

import (
	"encoding/base64"
	"net/http"
)

type INutanixClient interface {
	Host() string
	SetHost(host string)
	WithHost(host string) INutanixClient
	SetBasicAuth(username, password string)
}

type nutanixClient struct {
	host       string
	httpClient *http.Client
	authHeader http.Header
}

func (c *nutanixClient) Host() string {
	return c.host
}

func (c *nutanixClient) SetHost(host string) {
	c.host = host
}

func (c *nutanixClient) WithHost(host string) INutanixClient {
	c.host = host
	return c
}

func (c *nutanixClient) SetBasicAuth(username, password string) {
	c.authHeader = make(http.Header, 1)
	c.authHeader.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)),
	)
}

func (c *nutanixClient) GetAuthHeader() http.Header {
	return c.authHeader
}

func NewClient() INutanixClient {
	return &nutanixClient{
		httpClient: &http.Client{},
	}
}
