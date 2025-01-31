package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq" //go get github.com/lib/pq
)

const (
	HOST     = "localhost"
	PORT     = ":8055"
	USER     = "boss"
	PASSWORD = "password"
	DBNAME   = "boss_playground.db"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	RoleNo int    `json:"roleNo"`
	Phone  string `json:"phone"`
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	// connStr ="" //this also work but find why use above

	db, err := sql.Open("postgres", connStr) //it return database  or the err
	/*
		Errror to opening the Database  sql: unknown driver "postgres" (forgotten import?)
		this error ocurs when we not install the driver for postgress
	*/
	if err != nil {
		fmt.Println("Errror to opening the Database ", err)
		return
	}
	defer db.Close() //after  complint main last call to the databse clossing
	fmt.Println("Database Opening successfully")
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Starting the server")
	})

	router.HandleFunc("GET /api/users ", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Geting user")

	})

	router.HandleFunc("POST /api/create/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Creating the new User")
		var userData User

		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if errors.Is(err, io.EOF) {
				fmt.Fprintf(w, "Required body")
				return
			}
			fmt.Println("Error to conveting the struct", err)
			fmt.Fprint(w, "Getting the error")
			return
		}
		var id int
		row := db.QueryRow(
			`INSERT INTO user (name ,roleNo ,age ,phone,email)
			VALUES ( $1 , $2 ,$3 ,$4 ,$s)
			RETURNING id
			`,
			userData.Name, userData.RoleNo, userData.Age, userData.Phone, userData.Email,
		)
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("Error to inseting in db")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "error to Strore the user in database")
			return
		}

		fmt.Println("User inserted successfully with id: ", id)
		w.Header().Set("Content-Type", "appliction/json")
		w.WriteHeader(http.StatusCreated)

		respons := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "User Created Successfully",
		}

		json.NewEncoder(w).Encode(respons)

	})

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Println("Starting the server at port ", PORT)
	err = server.ListenAndServe()

	if err != nil {
		fmt.Println("Error to start the server", err)
	}

}
