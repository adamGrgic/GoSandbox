package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	fs := http.FS(staticFiles)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
