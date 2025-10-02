package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/fetch-formats", FetchFormatsHandler)
	http.HandleFunc("/api/download", DownloadHandler)

	// Health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("TeraBox Downloader API"))
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe":8080", nil))
}  
