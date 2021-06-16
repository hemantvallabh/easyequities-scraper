package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

func Login(username string, password string) (cookies map[string]string, err error) {

	c := getCollyInstance()
	statusCode := http.StatusInternalServerError
	loginErrorMessage := ""

	c.OnError(func(response *colly.Response, collyError error) {
		err = collyError
	})

	c.OnResponse(func(response *colly.Response) {
		statusCode = response.StatusCode
	})

	c.OnXML("/html/body/div[2]/div[2]/div/div[1]/div", func(e *colly.XMLElement) {
		loginErrorMessage = e.Text
	})

	// Build URL
	url, err := urlBuilder("/Account/SignIn", nil)
	if err != nil {
		return nil, err
	}

	// Post request
	err = c.Post(
		url,
		map[string]string{
			"UserIdentifier":  username,
			"Password":        password,
			"ReturnUrl":       "",
			"OneSignalGameId": "",
			"IsUsingNewLayoutSatrixOrEasyEquitiesMobileApp": "false",
		})
	if err != nil {
		return nil, err
	}

	c.Wait()

	// Handle successful response
	if statusCode == http.StatusOK && loginErrorMessage == "" {
		cookies = make(map[string]string, 0)
		for _, cookie := range c.Cookies(baseUrl) {
			cookies[cookie.Name] = cookie.Value
		}
		return cookies, nil
	}

	// Handle failure responses
	if statusCode != http.StatusOK {
		err = fmt.Errorf("http status code from server %d", statusCode)
	}

	if loginErrorMessage != "" {
		err = fmt.Errorf(loginErrorMessage)
	}

	return nil, err
}
