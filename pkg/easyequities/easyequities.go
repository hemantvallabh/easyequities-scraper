package easyequities

func Authentication(username string, password string) (string, error) {

	cookies, err := login(username, password)
	if err != nil {
		return "", err
	}

	authToken, err := encodeToAuthToken(cookies)
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func Accounts(authToken string) ([]Account, error) {

	accounts := make([]Account, 0)

	cookies, err := decodeFromAuthToken(authToken)
	if err != nil {
		return nil, err
	}

	accountIdentifiers, err := equityPage(cookies)
	if err != nil {
		return nil, err
	}

	for _, identifier := range accountIdentifiers {

		// Check currency/identifier available
		canUse, err := checkCurrencyAvailable(cookies, identifier.CurrencyId)
		if err != nil {
			return nil, err
		}
		if !canUse {
			continue
		}

		// Select identifier
		if err := selectAccount(cookies, identifier.AccountId); err != nil {
			return nil, err
		}

		accountValuation, err := accountOverviewPage(cookies, identifier.AccountId)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, Account{
			AccountToken:    encodeToAccountToken(identifier),
			AccountName:     accountValuation.TopSummary.AccountName,
			AccountNumber:   accountValuation.TopSummary.AccountNumber,
			AccountValue:    accountValuation.TopSummary.AccountValue,
			AccountCurrency: accountValuation.TopSummary.AccountCurrency,
			Movements: struct {
				ProfitLossValue      string
				ProfitLossPercentage string
			}{
				ProfitLossValue:      accountValuation.TopSummary.PeriodMovements[0].PercentageMove,
				ProfitLossPercentage: accountValuation.TopSummary.PeriodMovements[0].ValueMove,
			},
			Funds: struct {
				AvailableToInvest   string
				AvailableToWithdraw string
				UnsettledCash       string
				LockedFunds         string
			}{
				AvailableToInvest:   str(extractValue(fundSummaryAvailable, accountValuation.FundSummaryItems)),
				AvailableToWithdraw: str(extractValue(fundSummaryWithdrawal, accountValuation.FundSummaryItems)),
				UnsettledCash:       str(extractValue(fundSummaryUnsettledCash, accountValuation.FundSummaryItems)),
				LockedFunds:         str(extractValue(fundSummaryLockedFunds, accountValuation.FundSummaryItems)),
			},
		})
	}
	return accounts, nil
}
