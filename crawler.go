package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchURL(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func parseResponse(data string) {
	fmt.Println(data)
}
