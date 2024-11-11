package main

import (
	"os"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/mhson281/web-scraper/scraper"
)

func main() {
	godotenv.Load()

	baseURL := os.Getenv("BASE_URL")
	s, err := scraper.NewScraper(baseURL)
	if err != nil {
		log.Println("Error initializing scraper: ", err)
	}

	start := time.Now()
	s.Run()
	log.Printf("Scraping completed in %v\n", time.Since(start))
}
