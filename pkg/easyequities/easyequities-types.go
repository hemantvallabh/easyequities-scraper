package easyequities

type Account struct {
	AccountToken    string
	AccountName     string
	AccountNumber   string
	AccountValue    float64
	AccountCurrency string
	Movements       struct {
		ProfitLossValue      string
		ProfitLossPercentage string
	}
	Funds struct {
		AvailableToInvest   string
		AvailableToWithdraw string
		UnsettledCash       string
		LockedFunds         string
	}
}
