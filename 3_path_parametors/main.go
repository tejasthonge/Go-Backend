package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

/*
Feature	             Path Parameter (/user/{id})      	Query Parameter (/user?id=123)

Usage	             /user/123	                        /user?id=123
Best for	         Identifyingresources uniquely     	Filtering or modifying the request
Required?            Usually mandatory	                Optional, can have defaults
RESTful	             More RESTful,clean URLs	        Often used for filtering or optional data
*/

func main() {

	router := http.NewServeMux()
	//path parametors
	router.HandleFunc("GET /api/get/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		//by this way we get the path parametors

		id := r.PathValue("id")

		idInt, err := strconv.ParseInt(id, 0, 32)
		if err != nil {
			fmt.Fprintf(w, "Must pass valid id")
			return
		}
		fmt.Println(idInt)
		w.Write([]byte(id))
	})

	server := http.Server{
		Addr:    "localhost:8055",
		Handler: router,
	}

	slog.Info("Starting the server ...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("getting error at time of starting the server ", err)
	}

}
