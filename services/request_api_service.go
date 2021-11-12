package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/EdwinGamboaTotena/hackathonCeiba2021/models"
)

var (
	service *requestService
	lock    = &sync.Mutex{}
)

type requestService struct {
	host string
	port string
}

func GetInstance() *requestService {
	lock.Lock()
	defer lock.Unlock()

	if service == nil {
		host, _ := os.LookupEnv("SVC_API_HOSTNAME")
		port, _ := os.LookupEnv("SVC_API_PORT")
		service = &requestService{
			host: host,
			port: port,
		}
	}

	return service
}

func (rs *requestService) RequestApi(number string) (*models.Response, error) {
	client := getClient()
	resp, err := client.R().SetQueryParams(map[string]string{
		"number": number,
	}).SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("http://%s:%s/", rs.host, rs.port))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	response := models.Response{}
	err = json.Unmarshal(resp.Body(), &response)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &response, nil
}
