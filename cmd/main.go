package main

import (
	"github.com/hemantvallabh/easyequities-scraper/pkg/easyequities"
	"log"
	"os"
)

func main() {
	log.Println("Starting up....")

	authToken, err := easyequities.Authentication(os.Getenv("EE_UID"), os.Getenv("EE_PID"))
	if err != nil {
		log.Fatal(err)
	}

	if err = easyequities.Accounts(authToken); err != nil {
		log.Fatal(err)
	}
}
