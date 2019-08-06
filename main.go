package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting crawler...")

	url := "https://buyersguide.macrumors.com/#iOS"
	fmt.Println("Crawling:", url)

	var data = fetchURL(url)
	parseResponse(data)
}
