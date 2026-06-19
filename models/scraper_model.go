package models

type ScrapeRequest struct {
	URL string `json:"url"`
}

type ScrapeResult struct {
	TargetURL string `json:"target_url"`

	Links []string `json:"links"`

	Images []ImageResult `json:"images"`

	PageData PageData `json:"page_data"`
}


type ImageResult struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

// PageData holds page-level metadata pulled from the scraped document,
// separate from the raw links/images lists.
type PageData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`

	Headings []string `json:"headings"`

	LinkCount  int `json:"link_count"`
	ImageCount int `json:"image_count"`
}
