package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	greeting := "Hello, PipeCD "
	w.Write([]byte(greeting + version))
}

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/", handler)
	fmt.Println("Running...2")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Running...3")
}
