package easyequities

import (
	"encoding/base64"
)

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

func Logout(authToken string) error {

	cookies, err := decodeFromAuthToken(authToken)
	if err != nil {
		return err
	}

	return logout(cookies)
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

func Documents(authToken string) ([]Document, error) {

	documents := make([]Document, 0)

	cookies, err := decodeFromAuthToken(authToken)
	if err != nil {
		return nil, err
	}

	scrapedDocs, err := statementPage(cookies)
	if err != nil {
		return nil, err
	}

	for _, doc := range scrapedDocs {

		date, err := extractDocumentDate(doc.Description)
		if err != nil {
			continue
		}

		_, _, accountNumber, err := decodeDocumentUrl(doc.Url)
		if err != nil {
			continue
		}

		documents = append(documents, Document{
			Date:          date,
			Type:          doc.Type,
			DocumentToken: base64.StdEncoding.EncodeToString([]byte(doc.Url)),
			AccountNumber: accountNumber,
			FileName:      doc.FileName,
		})
	}

	return documents, nil
}

func DownloadDocument(authToken string, documentToken string) ([]byte, error) {

	cookies, err := decodeFromAuthToken(authToken)
	if err != nil {
		return nil, err
	}

	decodedUrl, err := base64.StdEncoding.DecodeString(documentToken)
	if err != nil {
		return nil, err
	}

	file, err := downloadStatement(cookies, string(decodedUrl))
	if err != nil {
		return nil, err
	}

	return file, nil
}
