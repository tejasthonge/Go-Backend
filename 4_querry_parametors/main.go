package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

/*
Feature	             Path Parameter (/user/{id})      	Query Parameter (/user?id=123)

Usage	             /user/123	                        /user?id=123
Best for	         Identifyingresources uniquely     	Filtering or modifying the request
Required?            Usually mandatory	                Optional, can have defaults
RESTful	             More RESTful,clean URLs	        Often used for filtering or optional data
*/

//http://localhost:8080/api/user?sortBy=Id
// http://localhost:8055/api/users?sortBy=id

func main() {

	router := http.NewServeMux()

	//geting the querry parametors
	router.HandleFunc("GET /api/users", func(w http.ResponseWriter, r *http.Request) {

		sortBy := r.URL.Query().Get("sortBy")

		if sortBy == "" {
			fmt.Fprint(w, "Getteing user without sorting")
			return
		}
		fmt.Fprintf(w, "Sorting the user by :%s", sortBy)

	})

	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}
	slog.Info("Starting the server ..")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Getting the error to starting the server ", err)
	}

}
