package main

import (
	"encoding/json"
	"github.com/hemantvallabh/easyequities-scraper/pkg/easyequities"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	log.Println("Starting up....")

	authToken, err := easyequities.Authentication(os.Getenv("EE_UID"), os.Getenv("EE_PID"))
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := easyequities.Accounts(authToken)
	if err != nil {
		log.Fatal(err)
	}

	str, _ := json.Marshal(accounts)
	log.Println(string(str))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
