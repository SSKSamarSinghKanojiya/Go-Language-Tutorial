


package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURl  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURl))
	fmt.Println("Hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Hasher Data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Encoded Hash: ", hash)
	fmt.Println("Final String: ", hash[:8])
	// return "https://short.url/" + hash[:8]
	return hash[:8]
}

func createURL(OriginalURl string) string {
	shortURL := generateShortURL(OriginalURl)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURl:  OriginalURl,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("GET Methods")
	fmt.Fprintf(w, "Welcome, Hello World")
}
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortURL_ := createURL(data.URL)
	// fmt.Fprintf(w, shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL((id))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)

	}
	http.Redirect(w, r, url.OriginalURl, http.StatusFound)

}

func main() {
	//  go mod init url-shortner
	// fmt.Println("Starting URL Shortener...")
	// OriginalURl := "https://github.com/SSKSamarSinghKanojiya/"

	// fmt.Println("Shortened URL: ", generateShortURL(OriginalURl))

	// Register the handler function to handle all request to the root URL ("/")
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server on PORT 3000
	fmt.Println("Starting Server on PORT 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}


/*
{
    "url":"https://github.com/SSKSamarSinghKanojiya/"
}
*/