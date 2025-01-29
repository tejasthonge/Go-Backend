package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/json", func(w http.ResponseWriter, r *http.Request) {

		//in go  if we have to sed the resopnse as json we have to use the structure

		response := struct {
			Message string      `json:"message"` //writing the resonpe go said tha user lowercase
			Data    interface{} `json:"data"`
		}{
			Message: "Ok",
			Data: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{
				Name: "Tejas Thonge",
				Age:  23,
			},
		}
		//if we not set the Content-Type then it will give by defout text/plain; charset=utf-8
		//so we hust have to set as application/json as bellow
		w.Header().Set("Content-Type", "application/json")
		//duet to this above it work properly
		

		json.NewEncoder(w).Encode(response) //this method convert struct to josn and werite in the resopne
	})

	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}

	server.ListenAndServe()
}
