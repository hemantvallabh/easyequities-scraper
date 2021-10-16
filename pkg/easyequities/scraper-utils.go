package easyequities

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"net/url"
)

func getCollyInstance() *colly.Collector {
	return colly.NewCollector(func(collector *colly.Collector) {
		collector.AllowURLRevisit = true
		collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"
	})
}

func getCollyInstanceWithCookies(url string, cookies map[string]string) *colly.Collector {

	cookieMonster := make([]*http.Cookie, 0)
	for key, value := range cookies {
		cookieMonster = append(cookieMonster, &http.Cookie{
			Name:  key,
			Value: value,
		})
	}

	return colly.NewCollector(func(collector *colly.Collector) {
		collector.AllowURLRevisit = true
		collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"
		collector.SetCookies(url, cookieMonster)
	})
}

func urlBuilder(path string, queryParams map[string]string) (string, error) {

	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	query := u.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	u.Path = path
	u.RawQuery = query.Encode()

	return u.String(), nil
}

func handleScrapingError(collector *colly.Collector, response *scrapingResponse) {

	collector.OnError(func(collyResponse *colly.Response, err error) {
		response.err = err
	})
}

func handleScrapingResponse(collector *colly.Collector, response *scrapingResponse) {

	collector.OnResponse(func(collyResponse *colly.Response) {
		response.statusCode = collyResponse.StatusCode
		response.url = collyResponse.Request.URL.String()
		response.response = collyResponse
	})
}

func evaluateScrapingResponse(response *scrapingResponse, urlPath string) error {

	if response.err != nil {
		return response.err
	}

	if response.statusCode != http.StatusOK {
		return fmt.Errorf("http status code != OK while scraping equity page. Status code [%d]", response.statusCode)
	}

	if !response.scrapCompleted {
		return fmt.Errorf("scraping incomplete")
	}

	if urlPath != "" {

		responseUrl, _ := url.Parse(response.url)
		if responseUrl.Path != urlPath {
			return fmt.Errorf("expected urlPath wasn't visited. urlPath expected [%s] visited [%s]", urlPath, responseUrl.Path)
		}
	}

	return nil
}

func extractValue(key string, data []map[string]interface{}) interface{} {

	for _, m := range data {
		labelName, ok := m[labelKey]
		if ok {
			if labelName == key {
				v, ok := m[valueKey]
				if ok {
					return v
				}
			}
		}
	}
	return nil
}

func str(t interface{}) string {
	if t == nil {
		return ""
	}

	return t.(string)
}
