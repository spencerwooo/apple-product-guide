package main

import (
	"fmt"

	"github.com/robfig/cron"
)

const iammain string = "[Main]"
const port string = "8000"

func main() {
	initUpdateTask()

	fmt.Printf("%s Starting server on %s...\n", iammain, "localhost:"+port)
	serve()
}

func initUpdateTask() {
	// Create cron task to update data hourly
	c := cron.New()
	c.AddFunc("@hourly", updateData)
	c.Start()
}

func updateData() {
	fmt.Printf("%s Starting crawler...\n", iammain)

	url := "https://buyersguide.macrumors.com"

	// Get dom elements from url
	var doc = fetchURL(url)
	// Parse dom elements with goquery
	parseResponse(doc)

	fmt.Printf("%s Latest data updated.\n", iammain)
}
