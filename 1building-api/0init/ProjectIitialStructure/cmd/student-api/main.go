package main

import (
	"fmt"
	"net/http"

	"github.com/tejasthonge/Go-Backend/1building-api/ProjectIitialStructure/pkg/config"
)

func main() {
	fmt.Println("Jay Shree Ram,\n wellcome to student api")
	cfg := config.MustLoad()
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Jay Shree Ram"))
	})

	fmt.Printf("Server Started at fort %s", cfg.Addr)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	server.ListenAndServe()

}
