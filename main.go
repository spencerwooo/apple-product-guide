package main

import (
	"fmt"
	"os"
	"apple-product-guide/utils"

	"github.com/robfig/cron"
)

const iammain string = "[Main]"

// Port Read from Heroku app
var Port = os.Getenv("PORT")

func main() {
	// Development environment port bindings
	if Port == "" {
		Port = "8000"
	}

	// Set $PORT
	utils.Port = Port

	// Update data on first launch
	updateData()
	// Set up cron job to update data every hour
	initUpdateTask()

	fmt.Printf("%s Server starting on %s...\n", iammain, Port)
	utils.Serve()
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
	var doc = utils.FetchURL(url)
	// Parse dom elements with goquery
	utils.ParseResponse(doc)

	fmt.Printf("%s Latest data updated.\n", iammain)
}
