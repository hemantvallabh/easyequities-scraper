package main

import (
	"encoding/json"
	"github.com/hemantvallabh/easyequities-scraper/pkg/easyequities"
	"log"
	"os"
)

func main() {
	log.Println("Starting up....")

	// Authenticate
	authToken, err := easyequities.Authentication(os.Getenv("EE_UID"), os.Getenv("EE_PID"))
	if err != nil {
		log.Fatal(err)
	}

	// Get all accounts
	accounts, err := easyequities.Accounts(authToken)
	if err != nil {
		log.Fatal(err)
	}
	str, _ := json.Marshal(accounts)
	log.Println(string(str))

	// Get all available documents
	documents, err := easyequities.Documents(authToken)
	if err != nil {
		log.Fatal(err)
	}
	str, _ = json.Marshal(documents)
	log.Println(string(str))

	// Download document
	if len(documents) > 0 {
		file, err := easyequities.DownloadDocument(authToken, documents[0].DocumentToken)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("c:\\temp\\" + documents[0].FileName, file, 0644);
		if err != nil {
			log.Fatal(err)
		}
	}

	// todo: holding, per account
	// todo: all available stocks
	// todo: transaction history
}
