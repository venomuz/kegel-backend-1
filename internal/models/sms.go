package models

type EskizToken struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

const EskizBaseUrl = "https://notify.eskiz.uz/api"
