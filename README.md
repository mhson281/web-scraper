# Go Web Scraper for Dead Link Detection

This Go application is a web scraper designed to detect dead links on a website. A dead link is defined as one that returns a 4xx or 5xx HTTP status code. This application uses concurrent goroutines to speed up the process of link checking, making it efficient and scalable for large websites.

## Features

- **Recursive Crawling**: Checks every link on pages within the same domain.
- **External Link Detection**: Identifies links to external domains and checks their status without further recursion.
- **Dead Link Detection**: Flags links with 4xx and 5xx status codes as "dead" and records them.
- **Concurrency**: Each link is processed in its own goroutine to optimize speed.
- **Safe Concurrency Management**: Uses `sync.Mutex` and `sync.WaitGroup` to handle concurrent access to shared resources.

## Requirements

- **Go** (version 1.16 or higher recommended)
- **Dependencies**: `golang.org/x/net/html` for HTML parsing.

## Installation

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/yourusername/webscraper.git
cd webscraper
Install the required Go module dependencies:
```

```bash
go mod tidy
```

## Usage

To run the scraper, use the following command:

```bash
go run main.go
```

## Project Structure

```plaintext

webscraper/
├── main.go             # Entry point of the application
├── scraper/
│   ├── scraper.go      # Core scraping logic
│   ├── checker.go      # Link status checking logic
│   ├── parser.go       # HTML parsing for anchor tags
│   └── utils.go        # Utility functions (e.g., URL resolution)
└── go.mod              # Go module file

```

### Explanation of Modules

- main.go: Sets up the initial scraper and runs it on the specified URL.
- scraper/scraper.go: Contains the Scraper struct and main logic for crawling links within the same domain.
- scraper/checker.go: Provides the CheckLink function to verify if a link is "dead."
- scraper/parser.go: Extracts anchor tags (<a href="...">) from HTML content.
- scraper/utils.go: Contains utility functions, such as ResolveURL for converting relative links to absolute URLs.

### How It Works

- Initialize Scraper: The Scraper struct is initialized with a base URL and begins by processing that URL.
- Link Checking: For each link on the page:
    - If it’s an external link (a different domain), the scraper checks if it’s dead but does not follow it.
    - If it’s an internal link (same domain), the scraper recursively follows and checks all links within that page.
- Concurrency Management: Each link is checked in a separate goroutine. sync.WaitGroup ensures that the main program waits until all goroutines are complete before printing results. sync.Mutex is used to safely update shared resources (visited links and deadLinks).

Example Output

```plaintext
Dead links found:
https://scrape-me.dreamsofcode.io/broken-link1
https://scrape-me.dreamsofcode.io/broken-link2
```
Screenshot:
![Example Ouptut](./asset/web_scraper_output.png)

## Notes

- This scraper only works with static HTML content. It does not handle JavaScript-rendered pages.
- External links are checked for dead status but are not recursively crawled.
- Make sure to run this scraper responsibly, as it may generate significant traffic for large websites.
