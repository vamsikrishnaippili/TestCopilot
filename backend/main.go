package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/api/fetch-formats", FetchFormatsHandler)
    http.HandleFunc("/api/download", DownloadHandler)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Write([]byte("TeraBox Downloader API"))
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server started at :%s\n", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
