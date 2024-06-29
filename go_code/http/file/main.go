package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024)

	})
	http.ListenAndServe(":8080", nil)
}
