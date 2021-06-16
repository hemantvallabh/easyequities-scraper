package easyequities

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hemantvallabh/easyequities-scraper/pkg/scraper"
)

func Authentication (username string, password string) (string, error) {

	cookies, err := scraper.Login(username, password)
	if err != nil {
		return "", err
	}

	jsonString, err := json.Marshal(cookies)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonString), nil
}

func Accounts (authToken string) {

}