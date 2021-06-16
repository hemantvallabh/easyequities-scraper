package scraper

import (
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

	cookieMonster := make([]*http.Cookie, len(cookies))
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
