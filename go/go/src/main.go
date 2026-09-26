package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Healthy!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
