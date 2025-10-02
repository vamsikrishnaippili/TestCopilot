package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type FetchFormatsRequest struct {
	Link string `json:"link"`
}
type FetchFormatsResponse struct {
	Success bool     `json:"success"`
	Formats []string `json:"formats,omitempty"`
	Message string   `json:"message,omitempty"`
}

func FetchFormatsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req FetchFormatsRequest
	_ = json.NewDecoder(r.Body).Decode(&req)
	formats, err := FetchTeraBoxFormats(req.Link)
	if err != nil {
		json.NewEncoder(w).Encode(FetchFormatsResponse{Success: false, Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(FetchFormatsResponse{Success: true, Formats: formats})
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	link := r.URL.Query().Get("link")
	format := r.URL.Query().Get("format")
	reader, fileName, err := DownloadTeraBoxVideo(link, format)
	if err != nil {
		http.Error(w, "Download failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, reader)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}  
