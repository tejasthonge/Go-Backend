package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func initServeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is an intial rout : %s", r.URL.Path)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	//here we use the all the methods on the http.Requst struct

	//1 Method:
	//	- it will return the method present in Request
	// (GET, POST, etc.).
	fmt.Println("method : ", r.Method) //method :  GET

	// 2 URL
	// - The requested URL (*url.URL object).
	fmt.Println("url : ", r.URL) //url :  /api/users

	//3 r.Host	Returns the requested host.
	fmt.Println("host :", r.Host) //host : localhost:8055

	// 4 Returns the HTTP protocol version (e.g., "HTTP/1.1").
	fmt.Println("Prtocol: ", r.Proto) //Prtocol:  HTTP/1.1

	// 5  A map containing HTTP headers.
	// fmt.Println("Headrs : ", r.Header) //Headrs :  map[Accept:[*/*] Accept-Encoding:[gzip, deflate, br] Cache-Control:[no-cache] Connection:[keep-alive] Postman-Token:[647ce6a3-a6d0-42e4-a36a-0579b5907ecd] User-Agent:[PostmanRuntime/7.43.0]]

	//6 r.RemoteAddr	Returns the client's IP address and port.
	fmt.Println("address of remote matchion : ", r.RemoteAddr) //address of remote matchion :  127.0.0.1:49820

	w.Write([]byte("ok"))
}


func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println("id : ", id)
}
func postUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Stroring the user")

	//geting the users body
	body := r.Body
	byteBody, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("Gett error for reading the body")
		return
	}
	fmt.Println("Body :", string(byteBody))

	w.WriteHeader(http.StatusCreated)
	w.Write(byteBody)
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /", initServeHandler)

	router.HandleFunc("GET /api/users", getUsersHandler)

	router.HandleFunc("GET /api/user:id", getSingleUserHandler)

	router.HandleFunc("POST /api/user", postUserHandler)

	server := http.Server{
		Addr:    "localhost:8055",
		Handler: router,
	}

	slog.Info("Starting server at prot localhost:8055")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error for Starting server .. %v", err)
	}

}
