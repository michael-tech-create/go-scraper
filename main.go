package main

import (
	"fmt"
	"go-scraper/handlers"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("template"))

	http.Handle("/", fs)
	http.HandleFunc("/api/scrape", handlers.HandleScrape)
	fmt.Println("api is running on port http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
