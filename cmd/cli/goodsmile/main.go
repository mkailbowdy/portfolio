package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Flag Configurations
	path := flag.String("path", "en/product", "The path to the product")
	flag.Parse()
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.goodsmile.com"),
	)

	c.OnHTML("div.b-product-info__header", func(e *colly.HTMLElement) {
		name := e.ChildText("h1")
		fmt.Printf("%s\n============\n", name)
	})

	c.OnHTML("section.p-product__section", func(e *colly.HTMLElement) {
		desc := e.ChildText("p[name=description]")
		if desc != "" {
			fmt.Printf("%s\n============\n", desc)
		}
	})

	c.OnHTML("div#specification dl dd", func(e *colly.HTMLElement) {
		spec := e.ChildText("p")
		if spec != "" {
			fmt.Printf("%s\n\n", spec)
		}
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.goodsmile.com/" + *path)
}
