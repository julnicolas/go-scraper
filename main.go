package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
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

// CLI is a struct reprensenting accepted CLI arguments
type CLI struct {
	URL    string
	Output func(*colly.HTMLElement)
}

func (o *CLI) Parse() error {
	var output string

	flag.StringVar(&o.URL, "u", "", "Url to collect links from. Can be used several times")
	flag.StringVar(&output, "o", "stdout", "Output formats. Possible formats are 'stdout' and 'json'")
	flag.Parse()

	if o.URL == "" {
		return fmt.Errorf("cli error - URL argument is empty")
	}

	switch strings.ToLower(output) {
	case "stdout":
		o.Output = Stdout(o.URL)

	case "json":
		o.Output = Json(o.URL)

	default:
		return fmt.Errorf("cli error - unsupported format '%s'", output)
	}

	return nil
}

func main() {
	cli := &CLI{}
	if err := cli.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Instantiate default collector
	// There is a very useful whitelist option in case
	// it is decided to recursively crawl URLs
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", cli.Output)

	// Start scraping on https://hackerspaces.org
	if err := c.Visit(cli.URL); err != nil {
		fmt.Printf("request error - %s\n", err)
		os.Exit(1)
	}
}
