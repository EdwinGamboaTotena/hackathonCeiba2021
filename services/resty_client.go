package services

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

var singletonClient *resty.Client

const (
	RetryCount    = 5
	RetryWaitTime = 200
)

func getClient() *resty.Client {
	lock.Lock()
	defer lock.Unlock()

	if singletonClient == nil {
		singletonClient = resty.New().
			SetRetryCount(RetryCount).
			SetRetryWaitTime(RetryWaitTime * time.Millisecond).
			AddRetryCondition(
				func(r *resty.Response, err error) bool {
					return r.StatusCode() >= http.StatusInternalServerError
				})
	}

	return singletonClient
}
