package services

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"sync"
	"time"

	"github.com/EdwinGamboaTotena/hackathonCeiba2021/models"
)

const (
	ERROR          = "error"
	HTTP           = "http"
	SvcApiHostname = "SVC_API_HOSTNAME"
	SvcApiPort     = "SVC_API_PORT"
)

var (
	service *requestService
	lock    = &sync.Mutex{}
)

type requestService struct {
	host  string
	port  string
	cache *cache.Cache
}

func GetInstance() *requestService {
	lock.Lock()
	defer lock.Unlock()

	if service == nil {
		initHost, initPort := initVarEntorn()
		initCache := cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
		service = &requestService{
			host:  initHost,
			port:  initPort,
			cache: initCache,
		}
	}

	return service
}

func initVarEntorn() (string, string) {
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

func (requestService *requestService) RequestApi(number string) (*models.Response, error) {
	response := models.Response{}

	value, found := requestService.cache.Get(number)
	if found {
		response = value.(models.Response)
		return &response, nil
	}

	value, found = requestService.cache.Get(ERROR)
	if found {
		return nil, fmt.Errorf(value.(string))
	}

	err := requestService.doRequest(number, &response)
	if err != nil {
		requestService.cache.Set(ERROR, err.Error(), 2*time.Second)
		return nil, err
	}

	requestService.cache.Set(number, response, time.Duration(response.ValiditySeconds)*time.Second)
	return &response, nil
}

func (requestService *requestService) doRequest(number string, response *models.Response) error {
	client := getClient()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"number": number,
		}).
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s://%s:%s/", HTTP, requestService.host, requestService.port))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return json.Unmarshal(resp.Body(), &response)
}
