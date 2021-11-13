package services

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/sony/gobreaker"

	"github.com/EdwinGamboaTotena/hackathonCeiba2021/models"
)

const (
	ERROR           = "error"
	HTTP            = "http"
	SvcApiHostname  = "SVC_API_HOSTNAME"
	SvcApiPort      = "SVC_API_PORT"
	RetryCount      = 5
	RetryWaitTime   = 200
	CircuitOpenTime = 1
	Requests        = 3
	FailureRatio    = 0.9
)

var (
	service *requestService
	lock    = &sync.Mutex{}
)

type requestService struct {
	host           string
	port           string
	client         *resty.Client
	cache          *cache.Cache
	circuitBreaker *gobreaker.CircuitBreaker
}

func GetInstance() *requestService {
	lock.Lock()
	defer lock.Unlock()

	if service == nil {
		initHost, initPort := initVarEnvironment()
		service = &requestService{
			host:           initHost,
			port:           initPort,
			client:         initClient(),
			cache:          cache.New(cache.DefaultExpiration, cache.DefaultExpiration),
			circuitBreaker: initCircuit(),
		}
	}

	return service
}

func (requestService *requestService) RequestApi(number string) (*models.Response, error) {
	response := models.Response{}

	value, found := requestService.cache.Get(number)
	if found {
		response = value.(models.Response)
		return &response, nil
	}

	_, found = requestService.cache.Get(ERROR)
	if found {
		time.Sleep(RetryWaitTime * time.Millisecond)
	}

	err := requestService.doRequest(number, &response)
	if err != nil {
		requestService.cache.Set(ERROR, err.Error(), CircuitOpenTime*time.Second)
		return nil, err
	}

	requestService.cache.Set(number, response, time.Duration(response.ValiditySeconds)*time.Second)
	return &response, nil
}

func (requestService *requestService) doRequest(number string, response *models.Response) error {
	body, err := requestService.circuitBreaker.Execute(func() (interface{}, error) {
		resp, err := requestService.client.R().
			SetQueryParams(map[string]string{
				"number": number,
			}).
			SetHeader("Accept", "application/json").
			Get(fmt.Sprintf("%s://%s:%s/", HTTP, requestService.host, requestService.port))

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return resp.Body(), nil
	})

	if err != nil {
		return err
	}

	return json.Unmarshal(body.([]byte), &response)
}
