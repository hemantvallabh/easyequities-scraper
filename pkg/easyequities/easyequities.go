package easyequities


func Authentication (username string, password string) (string, error) {

	cookies, err := login(username, password)
	if err != nil {
		return "", err
	}

	authToken, err := encodeToAuthToken(cookies)
	if err != nil{
		return "", err
	}

	return authToken, nil
}

func Accounts (authToken string) error {

	cookies, err := decodeFromAuthToken(authToken)
	if err != nil {
		return err
	}

	accounts, err := equityPage(cookies)
	if err != nil {
		return err
	}

	for _, account := range accounts {

		// Check currency/account available
		canUse, err := checkCurrencyAvailable(cookies, account.currencyId)
		if err != nil {
			return err
		}
		if !canUse {
			continue
		}

		// Select account
		if err := selectAccount(cookies, account.accountId); err != nil {
			return err
		}


	}

	// loopy
		// Select account
		// Change account
		// Navigate to /AccountsOverview
		// Navigate where else ???

		return nil
}