package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var shortenedURLs = make(map[string]string)

type URL struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/add", add)
	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := URL{}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		log.Fatal("Error wilst unmarshaling json", err)
		return
	}
	// map["1k8sckod9"] = "google.com"
	shortenedURLs[url.Short] = url.Long
	fmt.Print(shortenedURLs)
	json.NewEncoder(w).Encode(shortenedURLs)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := URL{}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		log.Fatal("Error wilst unmarshaling json", err)
		return
	}
	// map["1k8sckod9"] = "google.com"
	fmt.Print(shortenedURLs)
	json.NewEncoder(w).Encode(shortenedURLs[url.Short])
}
