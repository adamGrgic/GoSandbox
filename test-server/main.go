package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Starting up GO test server")

	http.HandleFunc("/", helloHandler)

	port := os.Getenv("TS-1-PORT")

	fmt.Println("Starting server on " + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}

	// select {}
}
