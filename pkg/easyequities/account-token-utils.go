package easyequities

import (
	"encoding/base64"
	"encoding/json"
)

func encodeToAccountToken(identifier accountIdentifier) string {
	jsonObj, _ := json.Marshal(identifier)
	return base64.StdEncoding.EncodeToString(jsonObj)
}

func decodeFromAccountToken(accountToken string) (accountIdentifier, error) {

	jsonObj, err := base64.StdEncoding.DecodeString(accountToken)
	if err != nil {
		return accountIdentifier{}, err
	}

	var identifier accountIdentifier
	err = json.Unmarshal(jsonObj, &identifier)
	if err != nil {
		return accountIdentifier{}, err
	}

	return identifier, nil
}
