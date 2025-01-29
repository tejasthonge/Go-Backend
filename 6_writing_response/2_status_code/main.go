package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /api/statuscode", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusAccepted)              //tis const are store in http package
		w.Write([]byte("Successfully Accepted Requst")) //this have connet type is

	})

	server := http.Server{
		Addr:    ":8055",  //
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("geting error : ", err)
	}
}
