package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchURL(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func parseResponse(doc *goquery.Document) {
	doc.Find(".guide.ios").Each(func(i int, product *goquery.Selection) {
		// iPhone XR
		name := product.Find("h2").Text()

		// Caution
		advice := product.Find("strong").First().Text()

		// Approaching End of Cycle
		status := strings.TrimSpace(strings.TrimPrefix(product.Find(".status").Text(), advice))

		// Mar 2019
		releaseDate := strings.TrimSpace(product.Find(".date").First().Text())
		// 141
		daysSinceLastRelease := strings.TrimSpace(product.Find("span").First().Text())
		// 390
		average := strings.TrimSpace(product.Find(".right.average").Text())
		fmt.Println(i, name+":", advice+",", status, daysSinceLastRelease, releaseDate, average)
	})
}
