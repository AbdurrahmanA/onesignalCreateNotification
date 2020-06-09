package onesignal

import (
	"net/http"
	"sync"
)

var doOnce sync.Once

//Clientc NewClient
type Clientc struct {
	Client http.Client
}

var client *Clientc

//NewClient NewClient
func NewClient() *Clientc {
	if client == nil {
		doOnce.Do(func() {
			client = new(Clientc)
		})
	}
	return client
}
