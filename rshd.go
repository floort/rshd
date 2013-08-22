package main

import (
	"log"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(".",r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
