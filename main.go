package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	fmt.Printf("Hello, PipeCD %s\n", version)
}

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/", handler)
	fmt.Println("Running...2")
	http.ListenAndServe(":80", nil)
	fmt.Println("Running...3")
}
