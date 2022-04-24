package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", healthz)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	fmt.Fprintf(w, "ok")
	fmt.Fprintf(w, version)
}
