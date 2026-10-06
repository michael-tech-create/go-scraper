# Go-Scraper

Hi, I'm Michael-tech-create.

Go-Scraper is a Go web app that fetches a live page with [Colly](https://github.com/gocolly/colly) and shows links, images, and SEO metadata in a Tailwind dashboard.

## Live data

Each scrape is a fresh HTTP request. Nothing is cached.

1. The browser sends `POST /api/scrape` with `{ "url": "https://example.com" }`.
2. Colly visits that URL and parses the HTML it receives.
3. The API returns JSON. The dashboard renders it immediately.

There is no background poll or WebSocket. Submit the form again to fetch the page again.

JavaScript-rendered sites only return the initial HTML. Colly does not run a browser.

## Features

- One-shot scrape of a public URL from the dashboard or the API.
- SEO fields: title, meta description, meta keywords, and `h1` / `h2` headings.
- Every `a[href]`, resolved to an absolute URL, with a count.
- Image `src` and `alt`, resolved to absolute URLs, shown in a gallery.
- A random desktop user agent plus common browser headers on each visit.
- Rate limit: 2 parallel requests, 1 second delay, plus up to 2 seconds of jitter.
- CSV export of the active tab (links, images, or headings) and a full JSON download.
- Light and dark theme.

## Run

Requires Go 1.24 or newer.

```bash
git clone https://github.com/michael-tech-create/go-scraper.git
cd go-scraper
go run .
```

Open `http://localhost:8000`, paste a URL, and click **Scrape Target**.

On Render, leave the root directory empty. Build command: `go build -tags netgo -ldflags '-s -w' -o app`. Start command: `./app`. Render sets `PORT`; the server listens on that value and falls back to `8000` locally. The dashboard calls `/api/scrape` on the same host.

## API

`POST /api/scrape`

```bash
curl -s -X POST http://localhost:8000/api/scrape \
  -H 'Content-Type: application/json' \
  -d '{"url":"https://example.com"}'
```

```json
{
  "target_url": "https://example.com",
  "links": ["https://example.com/about"],
  "images": [{ "src": "https://example.com/logo.png", "alt": "Logo" }],
  "page_data": {
    "title": "Example Domain",
    "description": "",
    "keywords": "",
    "headings": ["Example Domain"],
    "link_count": 1,
    "image_count": 1
  }
}
```

## Layout

| Path | Role |
| --- | --- |
| `main.go` | Serves `template/` and `POST /api/scrape` on port 8000 |
| `api/scraper_handlers.go` | Colly collector and JSON response |
| `template/index.html` | Dashboard, theme toggle, tabs, CSV and JSON export |

## Limits

- One page per request. Links are listed and not followed.
- HTML only. No login, cookies, or JavaScript execution.
- Public pages only. A site can still refuse the request.

## Author

Michael-tech-create, software engineer.
