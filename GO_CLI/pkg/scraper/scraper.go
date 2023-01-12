package scraper

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

var c = colly.NewCollector(
	colly.AllowedDomains("www.github.com"),
	colly.AllowedDomains("github.com"),
)

func InitScraper(username string, url string) {
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("\nGetting the followers...")
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("\nSomething went wrong:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("\nHere are the king makers")
	})
	c.OnHTML("turbo-frame#user-profile-frame div.position-relative div.d-table div.col-9 a span.Link--primary", func(e *colly.HTMLElement) {
		fmt.Print("\n")
		fmt.Print(e.Text)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("\n\nThanks for using!")
	})
	c.Visit("https://github.com/" + username + "?tab=" + url)

}

func GetCollector() *colly.Collector {
	return c
}
