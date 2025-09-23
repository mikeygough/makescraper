package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type story struct {
	Title       string    `json:"title" selector:".titleline a"`
	Link        string    `json:"link"`
	ExtractedAt time.Time `json:"extracted_at"`
}

func main() {
	// slice to store stories
	var stories []story

	// instantiate default Collector
	c := colly.NewCollector()

	// grab title
	c.OnHTML(".athing:first-of-type", func(e *colly.HTMLElement) {
		s := &story{}
		e.Unmarshal(s)

		// add link
		s.Link = e.ChildAttr(".titleline a", "href")
		// add current timestamp
		s.ExtractedAt = time.Now()

		fmt.Printf("Title: %q\n", s.Title)
		fmt.Printf("Link: %q\n", s.Link)
		fmt.Printf("Extracted at: %v\n", s.ExtractedAt)

		stories = append(stories, *s)
	})

	// before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// start scraping here
	c.Visit("https://news.ycombinator.com/")

	// serialize to json
	jsonData, err := json.MarshalIndent(stories, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}

	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}
