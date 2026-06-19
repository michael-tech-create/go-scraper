package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ScrapeRequest struct {
	URL string `json:"url"`
}

type ImageResult struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type PageData struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	Headings    []string `json:"headings"`
	LinkCount   int      `json:"link_count"`
	ImageCount  int      `json:"image_count"`
}

type ScrapeResult struct {
	TargetURL string        `json:"target_url"`
	Links     []string      `json:"links"`
	Images    []ImageResult `json:"images"`
	PageData  PageData      `json:"page_data"`
}


var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
}


func HandleScrape(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req ScrapeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	result := ScrapeResult{
		TargetURL: req.URL,
		Links:     []string{},
		Images:    []ImageResult{},
		PageData: PageData{
			Headings: []string{},
		},
	}

	c := colly.NewCollector()



	c.Limit(&colly.LimitRule{
		DomainRegexp: `.*`,                     
		Parallelism:  2,                        
		Delay:        1 * time.Second,          
		RandomDelay:  2 * time.Second,          
	})


	c.OnRequest(func(r *colly.Request) {
	
		randomAgent := userAgents[rand.Intn(len(userAgents))]
		r.Headers.Set("User-Agent", randomAgent)
		
		// Inject standard headers that actual commercial browsers send by default
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")

		fmt.Printf("Engine disguised. Using identity: %s -> Scraping: %s\n", randomAgent[:40]+"...", r.URL.String())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "" {
			result.Links = append(result.Links, e.Request.AbsoluteURL(link))
		}
	})

	// Image Scraping Logic 
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src == "" {
			return
		}
		result.Images = append(result.Images, ImageResult{
			Src: e.Request.AbsoluteURL(src),
			Alt: e.Attr("alt"),
		})
	})

	// SEO Extractions
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

	err := c.Visit(req.URL)
	if err != nil {
		errorMsg := fmt.Sprintf(`{"error": "Target security system blocked access or domain offline: %v"}`, err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	result.PageData.LinkCount = len(result.Links)
	result.PageData.ImageCount = len(result.Images)

	json.NewEncoder(w).Encode(result)
}