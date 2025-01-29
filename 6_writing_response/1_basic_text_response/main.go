package main

import (
	"fmt"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /api/text", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(w, "This is Sample text response")
	})

	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}
	server.ListenAndServe()
}
