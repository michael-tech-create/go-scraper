# Hi am Michael-tech-create

# Go-Scraper Advanced Dashboard 🕷️
What started as a command-line backend text-processing tool has evolved into a fully interactive, production-ready web server. Go-Scraper is a high-performance, concurrent web scraping engine built with Go and the Colly framework, paired with a sleek, responsive frontend dashboard styled with Tailwind CSS.
This application exposes powerful scraping logic to the web, allowing users to extract SEO metadata, page hierarchies, asset links, and structural data in real-time.

## Features
Backend (Go + Colly)
- Concurrent Scraping: Leverages Colly's asynchronous capabilities with configurable parallelism and random delays to prevent server overloads.

- Stealth & Evasion: Implements randomized User-Agent rotation and standard commercial browser headers to disguise the engine's identity.

- Rich Data Extraction: Automatically pulls:

- Complete SEO Meta Diagnostics (Title, Description, Keywords).

- Page typography and structural hierarchies (H1, H2 tags).

- Absolute URLs for all anchor links.

- Image assets including source URLs and Alt-text attributes.

- RESTful API: Exposes a clean POST /api/scrape endpoint that returns highly structured JSON.

## Frontend (HTML5 + Vanilla JS + Tailwind CSS)
- Modern UI: A clean dashboard built with Tailwind CSS, featuring full Light/Dark mode support.

- Real-Time Data Visualization: Splits incoming data into organized tabs (Links Panel, Images Gallery, Typography Outline).

- Data Export Pipeline: One-click functionality to export active scrape intelligence into CSV or full JSON formats.

# 📁 Project Structure
````
Plaintext
├── main.go                  # Application entry point and server configuration
├── handlers/
│   └── scraper_handlers.go  # Core scraping logic, API routing, and Colly configurations
└── template/
    └── index.html           # Frontend dashboard interface
````
## 🛠️ Installation & Setup
````

1 Clone the repository and initialize the module:

Bash
git clone <your-repo-url>
cd go-scraper
go mod init go-scraper

Install dependencies:
This project relies on the Colly framework for the heavy lifting.

Bash
go get github.com/gocolly/colly/v2
Run the server:

Bash
go run main.go

`````
# API Reference
POST /api/scrape
Initiates a scraping job on the provided target URL.

Request Body:

JSON
{
  "url": "https://example.com"
}
Success Response:

JSON
{
  "target_url": "https://example.com",
  "links": ["https://example.com/about", "..."],
  "images": [
    {
      "src": "https://example.com/logo.png",
      "alt": "Company Logo"
    }
  ],
  "page_data": {
    "title": "Example Domain",
    "description": "This domain is for use in illustrative examples.",
    "keywords": "",
    "headings": ["Example Domain"],
    "link_count": 1,
    "image_count": 1
  }
}

## Author
Michael-tech-create software engineer
