package scraper

import (
	"net/url"
	"log"
)

func ResolveURL(base *url.URL, href string) *url.URL {
	parsedURL, err := url.Parse(href)
	if err != nil {
		log.Printf("Error parsing URL %s: %v", href, err)
		return base
	}
	return base.ResolveReference(parsedURL)
}
