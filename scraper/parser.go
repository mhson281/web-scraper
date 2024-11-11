package scraper

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func ParseLinks(link string) ([]string, error) {
	resp, err := http.Get(link)
	if err != nil {
		log.Printf("Error fetching page %s: %v", link, err)
		return nil, err
	}
	defer resp.Body.Close()

	var links []string
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
	}

	return links, nil
}
