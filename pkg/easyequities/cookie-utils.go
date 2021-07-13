package easyequities

import (
	"encoding/base64"
	"encoding/json"
)

func encodeToAuthToken(cookies map[string]string) (string, error) {
	jsonObj, err := json.Marshal(cookies)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonObj), nil
}

func decodeFromAuthToken(authToken string) (map[string]string, error) {

	jsonObj, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		return nil, err
	}

	var cookies map[string]string
	err = json.Unmarshal(jsonObj, &cookies)
	if err != nil {
		return nil, err
	}

	return cookies, nil
}
