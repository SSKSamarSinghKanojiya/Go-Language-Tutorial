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

// URL represents the structure of a shortened URL entry
type URL struct {
	ID           string    `json:"id"`
	OriginalURl  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// In-memory storage (acts like a database)
var urlDB = make(map[string]URL)

// Generate a short hash from original URL
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:8] // Use first 8 characters
}

// Create a shortened URL entry and store it
func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURl:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// Retrieve the original URL by short ID
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

// Default root route
func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go URL Shortener 🚀")
}

// Handle /shorten route to create short URLs
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createURL(data.URL)

	// Prepare and return JSON response
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Redirect handler using shortened ID
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURl, http.StatusFound)
}

func main() {
	// Setup all routes
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the server
	fmt.Println("Starting Server on PORT 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

/*
📦 Sample JSON to test POST /shorten:

{
    "url": "https://github.com/SSKSamarSinghKanojiya/"
}

🌐 Then visit:
http://localhost:3000/redirect/{shortID}
*/















// package main

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type URL struct {
// 	ID           string    `json:"id"`
// 	OriginalURl  string    `json:"original_url"`
// 	ShortURL     string    `json:"short_url"`
// 	CreationDate time.Time `json:"creation_date"`
// }

// var urlDB = make(map[string]URL)

// func generateShortURL(OriginalURl string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(OriginalURl))
// 	fmt.Println("Hasher: ", hasher)
// 	data := hasher.Sum(nil)
// 	fmt.Println("Hasher Data: ", data)
// 	hash := hex.EncodeToString(data)
// 	fmt.Println("Encoded Hash: ", hash)
// 	fmt.Println("Final String: ", hash[:8])
// 	// return "https://short.url/" + hash[:8]
// 	return hash[:8]
// }

// func createURL(OriginalURl string) string {
// 	shortURL := generateShortURL(OriginalURl)  //It Convert the original URL string to a byte
// 	id := shortURL
// 	urlDB[id] = URL{
// 		ID:           id,
// 		OriginalURl:  OriginalURl,
// 		ShortURL:     shortURL,
// 		CreationDate: time.Now(),
// 	}
// 	return shortURL
// }

// func getURL(id string) (URL, error) {
// 	url, ok := urlDB[id]  // Use the short URL as the ID for Simplicity
// 	if !ok {
// 		return URL{}, errors.New("URL not found")
// 	}
// 	return url, nil
// }

// func RootPageURL(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println("GET Methods")
// 	fmt.Fprintf(w, "Welcome, Hello World")
// }
// func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
// 	var data struct {
// 		URL string `json:"url"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}
// 	shortURL_ := createURL(data.URL)
// 	// fmt.Fprintf(w, shortURL)
// 	response := struct {
// 		ShortURL string `json:"short_url"`
// 	}{ShortURL: shortURL_}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path[len("/redirect/"):]
// 	url, err := getURL((id))
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusNotFound)

// 	}
// 	http.Redirect(w, r, url.OriginalURl, http.StatusFound)

// }

// func main() {
// 	//  go mod init url-shortner
// 	// fmt.Println("Starting URL Shortener...")
// 	// OriginalURl := "https://github.com/SSKSamarSinghKanojiya/"

// 	// fmt.Println("Shortened URL: ", generateShortURL(OriginalURl))

// 	// Register the handler function to handle all request to the root URL ("/")
// 	http.HandleFunc("/", RootPageURL)
// 	http.HandleFunc("/shorten", ShortURLHandler)
// 	http.HandleFunc("/redirect/", redirectURLHandler)

// 	// Start the HTTP server on PORT 3000
// 	fmt.Println("Starting Server on PORT 3000...")
// 	err := http.ListenAndServe(":3000", nil)
// 	if err != nil {
// 		fmt.Println("Error starting server: ", err)
// 	}
// }


// /*
// {
//     "url":"https://github.com/SSKSamarSinghKanojiya/"
// }
// */