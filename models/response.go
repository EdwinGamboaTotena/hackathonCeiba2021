package models

type Response struct {
	Hostname        string `json:"hostname"`
	Method          string `json:"method"`
	Url             string `json:"url"`
	Data            string `json:"data"`
	Date            string `json:"date"`
	ValiditySeconds int    `json:"validitySeconds"`
	Token           string `json:"token"`
}
