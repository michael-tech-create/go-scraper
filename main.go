package main

import (
	"fmt"
	"go-scraper/api"
	"log"
	"net/http"
	"os"
)

func main() {

	fs := http.FileServer(http.Dir("template"))

	http.Handle("/", fs)
	http.HandleFunc("/api/scrape", handlers.HandleScrape)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := ":" + port
	fmt.Println("api is running on http://localhost" + addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
