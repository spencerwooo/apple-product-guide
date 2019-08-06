package main

import (
	"fmt"
)

const iammain string = "[Main]"

func main() {
	fmt.Printf("%s Starting crawler...\n", iammain)

	url := "https://buyersguide.macrumors.com"

	// Get dom elements from url
	var doc = fetchURL(url)
	// Parse dom elements with goquery
	parseResponse(doc)

	fmt.Printf("%s Latest data updated.\n", iammain)
}
