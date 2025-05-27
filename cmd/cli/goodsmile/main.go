package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

func main() {
	/* Flags */
	lang := flag.String("lang", "en", "the language to use")
	count := flag.Int("count", 2, "the number of items to create")
	flag.Parse()
	// Instantiate default collector
	var products []byte

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.goodsmile.com"),
	)

	var ranOnce bool
	c.OnHTML("div.c-top-product-list__unit", func(e *colly.HTMLElement) {
		if ranOnce {
			return
		}
		ranOnce = true

		e.ForEachWithBreak("a.c-top-product-list__item[href]", func(i int, h *colly.HTMLElement) bool {
			fmt.Println(i, *count)
			if i >= *count {
				return false
			}
			link := h.Attr("href")
			c.Visit(h.Request.AbsoluteURL(link))
			return true
		})
	})

	getDetails(c, "h1.b-product-info__title", products)
	getDetails(c, "span.c-price__main", products)
	getDetails(c, "p.b-product-info__note", products)
	getDetails(c, "p[name]", products)

	//c.OnHTML("h1.b-product-info__title", func(e *colly.HTMLElement) {
	//	title := e.DOM.Text()
	//	fmt.Printf("Title: %s\n", title)
	//})

	//c.OnHTML("span.c-price__main", func(e *colly.HTMLElement) {
	//	price := e.DOM.Text()
	//	fmt.Printf("price: %s\n", price)
	//})
	//c.OnHTML("p.b-product-info__note", func(e *colly.HTMLElement) {
	//	deliveryDate := e.DOM.Text()
	//	fmt.Printf("Delivery Date: %s\n", deliveryDate)
	//})
	//c.OnHTML("p[name]", func(e *colly.HTMLElement) {
	//	description := e.DOM.Text()
	//	description = strings.Split(description, ".")[0]
	//	fmt.Printf("Description: %s\n", description)
	//})

	c.OnHTML("div#specification", func(e *colly.HTMLElement) {
		e.ForEach("dl.b-outline-table__detail", func(_ int, dl *colly.HTMLElement) {
			term := dl.ChildText("dt h3")
			if term == "仕様" || term == "Specifications" {
				specText := dl.ChildText("dd p")
				fmt.Printf("Specification:%s\n\n\n", specText)
				fmt.Print("============================================\n\n")
			}
		})
	})

	c.Visit("https://www.goodsmile.com/" + *lang)
	os.WriteFile("goodsmile.txt", products, 0644)
}

func getDetails(c *colly.Collector, goquerySelector string, product []byte) {
	c.OnHTML(goquerySelector, func(e *colly.HTMLElement) {
		detail := e.DOM.Text()
		fmt.Printf("%s\n", detail)
		product = append(product, detail...)
	})
}
