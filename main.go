package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func Stdout(url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != url {
			fmt.Println(e.Request.AbsoluteURL(link))
		}
	}
}

func Json(url string) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		// link is useful for the hash section for the json version
		if link != url {
			fmt.Println(e.Request.AbsoluteURL(link))
		}
	}
}

func main() {
	const url string = "https://news.ycombinator.com"

	// Instantiate default collector
	// There is a very useful whitelist option in case
	// it is decided to recursively crawl URLs
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", Stdout(url))

	// Start scraping on https://hackerspaces.org
	c.Visit(url)
}
