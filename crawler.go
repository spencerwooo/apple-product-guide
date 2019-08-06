package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const iamcrawler string = "<Crawler>"

func fetchURL(url string) *goquery.Document {
	fmt.Printf("%s Crawling: %s\n", iamcrawler, url)

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

func parseDOM(doc *goquery.Document, element string, productList map[string]map[string]string) map[string]map[string]string {
	doc.Find(element).Each(func(i int, product *goquery.Selection) {
		productInfo := make(map[string]string)
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

		// productInfo["name"] = name
		productInfo["advice"] = advice
		productInfo["status"] = status
		productInfo["releaseDate"] = releaseDate
		productInfo["daysSinceLastRelease"] = daysSinceLastRelease
		productInfo["average"] = average

		// fmt.Println(productInfo)
		productList[name] = productInfo
	})

	return productList
}

func parseResponse(doc *goquery.Document) {
	products := make(map[string]map[string]string)

	// iOS lineups
	products = parseDOM(doc, ".guide.ios", products)

	// Mac lineups
	products = parseDOM(doc, ".guide.mac", products)

	// Others
	products = parseDOM(doc, ".guide.other", products)

	productsJSON, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	filePath := "data.json"
	err = ioutil.WriteFile(filePath, productsJSON, 0644)
	if err != nil {
		log.Fatal((err))
	} else {
		fmt.Printf("%s Saved data to %s\n", iamcrawler, filePath)
	}
}
