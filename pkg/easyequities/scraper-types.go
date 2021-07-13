package easyequities

import "github.com/gocolly/colly"

type scrapingResponse struct {
	err            error
	statusCode     int
	url            string
	scrapCompleted bool
	response       *colly.Response
}

type account struct {
	accountId   string
	currencyId  string
	description string
}

type canUseSelectedAccount struct {
	CanUse  bool   `json: CanUse`
	Message string `json: Message`
}

type accountOverview struct {
	accountNumber        string
	accountValue         string
	profitLossValue      string
	profitLossPercentage string
}
