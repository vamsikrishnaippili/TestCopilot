package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// Mock: Return available formats based on link
func FetchTeraBoxFormats(link string) ([]string, error) {
	if !strings.Contains(link, "terabox") {
		return nil, errors.New("Invalid TeraBox link")
	}
	return []string{"mp4", "mkv", "webm"}, nil
}

// Mock: Return a fake file stream for demonstration
func DownloadTeraBoxVideo(link, format string) (io.Reader, string, error) {
	if !strings.Contains(link, "terabox") {
		return nil, "", errors.New("Invalid TeraBox link")
	}
	// For demo, return a text file as the "video"
	content := []byte("This is a mock video file for " + link + " in format " + format)
	return bytes.NewReader(content), "video." + format, nil
}  
