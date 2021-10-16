package easyequities

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
)

func extractDocumentDate(date string) (time.Time, error) {

	s := strings.TrimSpace(strings.ReplaceAll(date, "\n", ""))

	if len(s) < 7 {
		return time.Now(), fmt.Errorf("incorrect date format")
	}

	return time.Parse("2006_01", s[:7])
}

func decodeDocumentUrl(urlString string) (urlPath string, documentId string, accountNumber string, err error) {

	u, err := url.Parse(urlString)
	if err != nil {
		return "", "", "", err
	}

	accountNumber = u.Query().Get("accountNumber")
	if accountNumber == "" {
		return "", "", "", fmt.Errorf("no account number present")
	}

	urlPath = u.Path
	documentId = path.Base(urlPath)

	return urlPath, documentId, accountNumber, err
}
