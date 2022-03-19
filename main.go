package main

import (
  "os"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  version := os.Getenv("VERSION")
	fmt.Printf("Hello, PipeCD %s\n", version)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
