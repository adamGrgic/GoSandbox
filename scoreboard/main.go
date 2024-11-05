package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	fs := http.FS(staticFiles)
	http.Handle("/", http.FileServer(fs))
	http.ListenAndServe(":8080", nil)
}
