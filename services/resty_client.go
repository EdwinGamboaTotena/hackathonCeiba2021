package services

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

var singletonClient *resty.Client

func getClient() *resty.Client {
	lock.Lock()
	defer lock.Unlock()

	if singletonClient == nil {
		singletonClient = resty.New().
			SetRetryCount(3).
			SetRetryWaitTime(1 * time.Second).
			AddRetryCondition(
				func(r *resty.Response, err error) bool {
					return r.StatusCode() >= http.StatusInternalServerError
				})
	}

	return singletonClient
}
