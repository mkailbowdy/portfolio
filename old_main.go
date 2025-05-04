package main

//
//import (
//	"fmt"
//
//	"github.com/gocolly/colly"
//)
//
//func main() {
//	// Instantiate default collector
//	c := colly.NewCollector(
//		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
//		colly.AllowedDomains("www.tokyodev.com"),
//	)
//
//	// On every a element which has href attribute call callback
//	c.OnHTML("ul li section h3 a[href]", func(e *colly.HTMLElement) {
//		link := e.Attr("href")
//		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
//		e.Request.Visit("https://www.tokyodev.com/" + link)
//	})
//
//	c.OnHTML("section", func(e *colly.HTMLElement) {
//		cn := e.ChildText("h1")
//		if cn != "" {
//			fmt.Printf("Title: %s\n", cn)
//		}
//		hl := e.ChildText("h2.headline")
//		if hl != "" {
//			fmt.Printf("Headline: %s\n\n", hl)
//		}
//	})
//
//	// Start scraping on https://hackerspaces.org
//	c.Visit("https://www.tokyodev.com/jobs/no-japanese-required")
//}
