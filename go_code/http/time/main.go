package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(now)
		fmt.Fprintf(w, now)

	})
	http.ListenAndServe(":8080", nil)
}
