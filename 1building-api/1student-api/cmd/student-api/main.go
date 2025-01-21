package main

import (
	"fmt"
	"net/http"

	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/config"
)

// go run cmd/student-api/main.go --config config/local.yaml
func main() {
	fmt.Println("Jay Shree Ram,\n wellcome to student api")
	cfg := config.MustLoad() // in this struct we have all the cofig varible
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {

		res.Write([]byte("Jay Shree Ram"))
	})

	fmt.Printf("Server Started at fort %s", cfg.Addr)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	server.ListenAndServe()

}
