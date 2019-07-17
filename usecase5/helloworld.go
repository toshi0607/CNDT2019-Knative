package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "NOT SPECIFIED"
	}
	fmt.Fprintf(w, "Hello World: %s!\n", target)
}

func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
