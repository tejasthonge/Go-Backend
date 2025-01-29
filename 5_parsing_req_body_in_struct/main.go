package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	//in other languge we create the model class and porse the
	//body in to the our model like
	//go have the struct

	router := http.NewServeMux()

	router.HandleFunc("POST /api/user", func(w http.ResponseWriter, r *http.Request) {
		var user User
		body := r.Body

		err := json.NewDecoder(body).Decode(&user) //in this line we copy all the body values in  users struct
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Fprint(w, "Getting The error to parsing the body in struct", err)

			return
		}

		fmt.Println("name ", user.Name)
		fmt.Println("age ", user.Age)
		userByt, err := json.Marshal(user) //
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Getting error to parsing struct in json", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(userByt)

	})

	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}
	slog.Info("Starting the Server ..")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Getting the error at starting the server ..", err)
	}
}
