package main
//this and basic eg
import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"


	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "boss"
	PASSWORD = "password"
	DBNAME   = "boss_playground" // do not use .db for postgres
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
	// Connect to PostgreSQL using the default postgres database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", HOST, PORT, USER, PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	defer db.Close() // Ensure the database connection is closed after main completes

	// Step 1: Try to create the new database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", DBNAME))
	if err != nil {
		if err.Error() == "pq: database \"boss_playground\" already exists" {
			fmt.Println("Database already exists:", DBNAME)
		} else {
			fmt.Println("Error creating database: ", err)
			return
		}
	} else {
		fmt.Println("Database created successfully with name:", DBNAME)
	}

	// Connect to the newly created database
	connStrNewDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	dbNew, err := sql.Open("postgres", connStrNewDB)
	if err != nil {
		fmt.Println("Error connecting to new database: ", err)
		return
	}
	defer dbNew.Close()

	// Step 2: Create the users table if it doesn't exist
	_, err = dbNew.Exec(`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER,
		roleNo INTEGER,
		phoneNo BIGINT
	)`)
	if err != nil {
		fmt.Println("Error creating the table:", err)
		return
	} else {
		fmt.Println("Table created successfully!")
	}

	// Step 3: Setting up HTTP routes
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Starting the server")
	})

	// Route to handle getting users (stub for now)
	router.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Getting users")
		// Here, you can fetch users from DB if needed
	})

	// Route to handle creating a new user
	router.HandleFunc("/api/create/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Creating new user")
		var userData User

		// Decode incoming JSON body
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if errors.Is(err, io.EOF) {
				fmt.Fprintf(w, "Required body")
				return
			}
			fmt.Println("Error decoding the request body", err)
			fmt.Fprint(w, "Error decoding the request body")
			return
		}

		// Insert user into the database
		var id int
		row := dbNew.QueryRow(
			`INSERT INTO users (name, roleNo, age, phoneNo, email)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`,
			userData.Name, userData.RoleNo, userData.Age, userData.Phone, userData.Email,
		)
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("Error inserting into DB:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error storing the user in database")
			return
		}

		// Send successful response
		fmt.Println("User inserted successfully with id:", id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		response := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "User created successfully",
		}

		// Return response as JSON
		json.NewEncoder(w).Encode(response)
	})

	// Step 4: Start the server
	server := http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	fmt.Println("Starting the server on port", PORT)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
