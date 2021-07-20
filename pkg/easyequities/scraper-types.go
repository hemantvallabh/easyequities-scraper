package easyequities

import "github.com/gocolly/colly"

type scrapingResponse struct {
	err            error
	statusCode     int
	url            string
	scrapCompleted bool
	response       *colly.Response
}

type accountIdentifier struct {
	AccountId  string
	CurrencyId string
}

type canUseSelectedAccount struct {
	CanUse  bool   `json: CanUse`
	Message string `json: Message`
}

type trustAccountValuation struct {
	TopSummary struct {
		AccountValue    float64
		AccountCurrency string
		AccountNumber   string
		AccountName     string
		PeriodMovements []struct {
			ValueMoveLabel      string
			ValueMove           string
			PercentageMoveLabel string
			PercentageMove      string
			PeriodMoveHeader    string
		}
	}
	FundSummaryItems           []map[string]interface{}
	NetInterestOnCashItems     []map[string]interface{}
	AccrualSummaryItems        []map[string]interface{}
	InvestmentTypesAndManagers struct {
		InvestmentTypes []map[string]interface{}
		Managers        []map[string]interface{}
	}
	InvestmentSummaryItems     []map[string]interface{}
	CostsSummaryItems          []map[string]interface{}
	AccrualIncomeSummaryItems  []map[string]interface{}
	AccrualExpenseSummaryItems []map[string]interface{}
}
