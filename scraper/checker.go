package scraper

import (
	"log"
	"net/http"
)

func CheckLink(link string) bool {
	resp, err := http.Get(link)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", link, err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		log.Printf("Dead link found: %s, StatusCode: %d", link, resp.Status)
		return true
	}
	return false
}
