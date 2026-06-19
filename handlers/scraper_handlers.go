package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-scraper/models"

	"github.com/gocolly/colly/v2"
)

func HandleScrape(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
		// early returns
	}

	var req models.ScrapeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var result models.ScrapeResult

	result = models.ScrapeResult{
		TargetURL: req.URL,
		Links:     []string{},
		Images:    []models.ImageResult{},
		PageData: models.PageData{
			Headings: []string{},
		},
	}

	c := colly.NewCollector()

	// existing: link scraping
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		result.Links = append(result.Links, link)
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src == "" {
			return
		}

		result.Images = append(result.Images, models.ImageResult{
			Src: e.Request.AbsoluteURL(src),
			Alt: e.Attr("alt"),
		})
	})


	c.OnHTML("title", func(e *colly.HTMLElement) {
		result.PageData.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		result.PageData.Description = e.Attr("content")
	})


	c.OnHTML(`meta[name="keywords"]`, func(e *colly.HTMLElement) {
		result.PageData.Keywords = e.Attr("content")
	})


	c.OnHTML("h1, h2", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		if text != "" {
			result.PageData.Headings = append(result.PageData.Headings, text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("scraping initiated: " + r.URL.String())
	})

	err := c.Visit(req.URL)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to scrape url: %v", err), http.StatusInternalServerError)
		return
	}

	result.PageData.LinkCount = len(result.Links)
	result.PageData.ImageCount = len(result.Images)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
