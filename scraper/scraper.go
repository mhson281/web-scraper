package scraper

import (
	"fmt"
	"net/url"
	"sync"
)

type Scraper struct {
	baseURL   *url.URL
	visited   map[string]bool
	deadLinks []string
	mu        sync.Mutex
	wg        sync.WaitGroup
}


// Create a new scraper from parsing URL
func NewScraper(baseURL string) (*Scraper, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Scraper{
		baseURL: parsedURL,
		visited: make(map[string]bool),
		deadLinks: make([]string, 0),
	}, nil
}

// Run the scraper
func (s *Scraper) Run() {
	s.wg.Add(1)
	go s.processURL(s.baseURL.String())
}

func (s *Scraper) processURL(link string) {
	defer s.wg.Done()

	s.mu.Lock()
	if s.visited[link] {
		s.mu.Unlock()
		return
	}
	s.visited[link] = true
	s.mu.Unlock()

	if CheckLink(link) {
		s.mu.Lock()
		s.deadLinks = append(s.deadLinks, link)
		s.mu.Unlock()
		return
	}

	links, err := ParseLinks(link)

}
