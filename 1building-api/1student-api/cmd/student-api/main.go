package main

import (
	"context"

	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/config"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/config/handlers/student"
)

// go run cmd/student-api/main.go --config config/local.yaml
func main() {
	slog.Info("Jay Shree Ram,\n wellcome to student api") //it is simler as the fmt.println
	cfg := config.MustLoad()                              // in this struct we have all the cofig varible
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {

		res.Write([]byte("Jay Shree Ram"))
	})

	router.HandleFunc("POST /api/student/create", student.New())

	slog.Info("Server Started at ", slog.String("Adress :", cfg.Addr))
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	//server.ListenAndServe()
	// this woring properly but
	//some time requste is prossing and due some resion server wase stoping that
	//time we have firs complite and running request and then we have to
	//stope the single
	//this is know as ## gressfullsutdown
	//it is bassicly stutting dowon serve after completing ongoing requist
	done := make(chan os.Signal, 1) //this is buffer chanal

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGABRT) //here we passing the chan and which type of signal is comminf for notify
	go func() {
		err := server.ListenAndServe() //if server is not able to start the it will returning the error otherwise it will returning the nil
		if err != nil {
			log.Fatal("Faild to Start Server !")
		}
	}()
	<-done

	slog.Info("Sutting Down the Server .. ")
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		slog.Error("Faild To Shutting down server", slog.String("Error :", err.Error()))
	}

	slog.Info("Serever Shutting down successfully !")
}
