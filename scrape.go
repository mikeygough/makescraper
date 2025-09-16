package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type story struct {
	Title string `selector:".titleline a"`
}

func main() {
	// instantiate default Collector
	c := colly.NewCollector()

	// grab title
	c.OnHTML(".athing:first-of-type", func(e *colly.HTMLElement) {
		s := &story{}
		e.Unmarshal(s)

		fmt.Printf("Title: %q\n", s.Title)
	})

	// before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// start scraping here
	c.Visit("https://news.ycombinator.com/")
}
