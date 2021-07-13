package easyequities

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strconv"
	"time"
)

func login(username string, password string) (cookies map[string]string, err error) {

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

func equityPage(cookies map[string]string) ([]account, error) {

	var scrapingResponse scrapingResponse
	var accounts = make([]account, 0)
	var urlPath = "/Equity"

	c := getCollyInstanceWithCookies(baseUrl, cookies)

	handleScrapingError(c, &scrapingResponse)
	handleScrapingResponse(c, &scrapingResponse)

	c.OnHTML("#trustaccount-slider", func(trustAccountSlider *colly.HTMLElement) {

		trustAccountSlider.ForEach("#trust-account-content", func(i int, trustAccount *colly.HTMLElement) {
			selectorTab := trustAccount.DOM.Find("#selector-tab")
			if selectorTab.Nodes != nil {

				account := account{}
				account.accountId, _ = selectorTab.Attr("data-id")
				account.currencyId, _ = selectorTab.Attr("data-tradingcurrencyid")
				account.description = selectorTab.Find("#trust-account-types").Text()

				accounts = append(accounts, account)
			}
		})

		scrapingResponse.scrapCompleted = true
	})

	// Build URL
	url, err := urlBuilder(urlPath, nil)
	if err != nil {
		return nil, err
	}

	// Visit the page
	if err = c.Visit(url); err != nil {
		return nil, err
	}
	c.Wait()

	// Check response
	if err := evaluateScrapingResponse(&scrapingResponse, urlPath); err != nil {
		return nil, err
	}

	return accounts, nil
}

func checkCurrencyAvailable(cookies map[string]string, currencyId string) (bool, error){

	var scrapingResponse scrapingResponse
	var urlPath = "/Menu/CanUseSelectedAccount"
	response := canUseSelectedAccount{}

	c := getCollyInstanceWithCookies(baseUrl, cookies)

	handleScrapingError(c, &scrapingResponse)
	handleScrapingResponse(c, &scrapingResponse)

	c.OnResponse(func(r *colly.Response) {

		if err := json.Unmarshal(r.Body, &response); err != nil {
			scrapingResponse.err = err
		}

		scrapingResponse.scrapCompleted = true
	})

	// Build URL
	url, err := urlBuilder(urlPath, map[string]string{
		"tradingCurrencyId": currencyId,
		"_": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil {
		return false, err
	}

	// Visit the page
	if err = c.Visit(url); err != nil {
		return false, err
	}
	c.Wait()

	// Check response
	if err := evaluateScrapingResponse(&scrapingResponse, urlPath); err != nil {
		return false, err
	}

	return response.CanUse, nil
}

func selectAccount(cookies map[string]string, accountId string) error {

	var scrapingResponse scrapingResponse
	var urlPath = "/Menu/UpdateCurrency"

	c := getCollyInstanceWithCookies(baseUrl, cookies)

	handleScrapingError(c, &scrapingResponse)
	handleScrapingResponse(c, &scrapingResponse)

	c.OnResponse(func(r *colly.Response) {
		scrapingResponse.scrapCompleted = true
	})

	// Build URL
	url, err := urlBuilder(urlPath, map[string]string{
		"trustAccountId": accountId,
	})
	if err != nil {
		return err
	}

	// Visit the page
	if err = c.Visit(url); err != nil {
		return err
	}
	c.Wait()

	// Check response
	if err := evaluateScrapingResponse(&scrapingResponse, urlPath); err != nil {
		return err
	}

	return nil
}

func accountOverviewPage (cookies map[string]string, accountId string) (accountOverview, error) {

	var overview accountOverview
	var scrapingResponse scrapingResponse
	var urlPath = "/Menu/UpdateCurrency"

	c := getCollyInstanceWithCookies(baseUrl, cookies)

	handleScrapingError(c, &scrapingResponse)
	handleScrapingResponse(c, &scrapingResponse)

	c.OnHTML("body > div.container.container-margin.container-min-height > div.mobile-margin-top-15 > div.text-center.holdings-slider-header.control-margin-bottom-large.account-overview-heading > div > h3 > span", func(element *colly.HTMLElement) {
		overview.accountNumber = element.Text
	})

	// Build URL
	url, err := urlBuilder(urlPath, map[string]string{
		"trustAccountId": accountId,
	})
	if err != nil {
		return accountOverview{}, err
	}

	// Visit the page
	if err = c.Visit(url); err != nil {
		return accountOverview{}, err
	}
	c.Wait()

	// Check response
	if err := evaluateScrapingResponse(&scrapingResponse, urlPath); err != nil {
		return accountOverview{}, err
	}

	return overview, nil
}