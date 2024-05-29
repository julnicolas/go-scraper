package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func Stdout(e *colly.HTMLElement) {
	link := e.Attr("href")
	fmt.Println(e.Request.AbsoluteURL(link))
}

func Json(e *colly.HTMLElement) {
	link := e.Attr("href")
	// Print link
	// link is useful for the hash section for the json version
	fmt.Println(e.Request.AbsoluteURL(link))
}

func main() {
	// Instantiate default collector
	// There is a very useful whitelist option in case
	// it is decided to recursively crawl URLs
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", Stdout)

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com")
}
