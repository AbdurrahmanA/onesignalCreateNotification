package onesignal

import (
	"net/http"
	"sync"
)

var doOnce sync.Once

//Client NewClient
type Client struct {
	Client http.Client
}

var client *Client

//NewClient NewClient
func (c Client) NewClient() *Client {
	if client == nil {
		doOnce.Do(func() {
			client = new(Client)
		})
	}
	return client
}
