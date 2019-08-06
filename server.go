package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const iamserver = "<Server>"

func serve() {
	r := mux.NewRouter()
	r.HandleFunc("/api", productListHandler)
	r.HandleFunc("/api/{product}", productDetailHandler)
	http.Handle("/", r)

	fmt.Printf("%s Server listening on port %s.\n", iamserver, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getData() map[string]map[string]string {
	// Read data.json first
	dataFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dataFile.Close()

	dataByte, _ := ioutil.ReadAll(dataFile)

	var data map[string]map[string]string
	json.Unmarshal(dataByte, &data)

	return data
}

func productListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s URL: %s\n", iamserver, r.URL.Path)

	data := getData()

	// Return product list lineup
	resp := make([]string, 0, len(data))
	for key := range data {
		resp = append(resp, key)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s URL: %s\n", iamserver, r.URL.Path)

	data := getData()
	params := mux.Vars(r)
	// Return product detail
	resp := data[params["product"]]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
