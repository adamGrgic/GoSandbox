package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Construct the path to your index.html file
	staticDir := filepath.Join(cwd, "/static")

	// Serve the static files
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	// Start the server
	http.ListenAndServe(":9500", nil)
}
