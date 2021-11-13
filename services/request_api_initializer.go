package services

import (
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sony/gobreaker"
)

func initVarEnvironment() (string, string) {
	initHost := "api-2.hack.local"
	initPort := "80"
	if val, ok := os.LookupEnv(SvcApiHostname); ok {
		initHost = val
	}
	if val, ok := os.LookupEnv(SvcApiPort); ok {
		initPort = val
	}
	return initHost, initPort
}

func initCircuit() *gobreaker.CircuitBreaker {
	var st gobreaker.Settings
	st.Name = "HTTP GET EXTERNAL API"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= Requests && failureRatio >= FailureRatio
	}

	return gobreaker.NewCircuitBreaker(st)
}

func initClient() *resty.Client {
	return resty.New().
		SetRetryCount(RetryCount).
		SetRetryWaitTime(RetryWaitTime * time.Millisecond).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() >= http.StatusInternalServerError
			})
}
