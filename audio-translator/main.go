package main

import (
	"audio-translator/processor"
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "Bad upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	outPath := "uploads/" + header.Filename
	outFile, err := processor.SaveFile(file, outPath)
	if err != nil {
		http.Error(w, "Could not save", http.StatusInternalServerError)
		return
	}

	go processor.Process(outFile) // non-blocking

	fmt.Fprintf(w, "Processing %s...\n", header.Filename)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
