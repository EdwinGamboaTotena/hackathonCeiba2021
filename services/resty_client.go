package services

import "github.com/go-resty/resty/v2"

var singletonClient *resty.Client

func getClient() *resty.Client {
	lock.Lock()
	defer lock.Unlock()

	if singletonClient == nil {
		singletonClient = resty.New()
	}

	return singletonClient
}
