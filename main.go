package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting crawler...")

	url := "https://buyersguide.macrumors.com"
	fmt.Println("Crawling:", url)

	// Get dom elements from url
	var doc = fetchURL(url)
	// Parse dom elements with goquery
	parseResponse(doc)
}
