package main

import (
	"fmt"
	"os"

	"github.com/robfig/cron"
)

const iammain string = "[Main]"

// Read $PORT from Heroku app
var port = os.Getenv("PORT")

func main() {
	initUpdateTask()

	// Development environment port bindings
	if port == "" {
		port = "8000"
	}

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
