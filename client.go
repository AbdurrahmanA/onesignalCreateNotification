package onesignal

import (
	"net/http"
	"sync"
)

var doOnce sync.Once

//Client ss
type Client struct {
	Client http.Client
}

var client *Client

//NewClient ss
func NewClient() *Client {
	if client == nil {
		doOnce.Do(func() {
			client = new(Client)
		})
	}
	return client
}
