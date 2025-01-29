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
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/handlers/student"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/storage/sqlite"
)

// go run cmd/student-api/main.go --config config/local.yaml
func main() {
	slog.Info("Jay Shree Ram,\n wellcome to student api") //it is simler as the fmt.println
	//Statep 1 :laoading all the confg
	cfg := config.MustLoad() // in this struct we have all the cofig varible

	//stape 2 : connecting the database
	sqliteStorage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	slog.Info("Storage Initialized successully", slog.String("Env: ", cfg.Env), slog.String("Storage path: ", cfg.StoragePath), slog.String("Version", "1.0.1"))

	router := http.NewServeMux()
	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {

		res.Write([]byte("Jay Shree Ram"))
	})

	router.HandleFunc("POST /api/student/create", student.New(sqliteStorage))

	slog.Info("Server Started at ", slog.String("Adress :", cfg.Addr))
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	//server.ListenAndServe()
	// this working properly but
	//some time requste is prossing and if due some resione server wase stoping that
	//time we have first complite and running request and then we have to
	//stope the single
	//this is know as ## gressfullsutdown
	//it is bassicly stutting dowon serve after completing ongoing request
	done := make(chan os.Signal, 1) //this is buffer chanal

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGABRT) //here we passing the chan and which type of signal is comminf for notify
	go func() {
		err := server.ListenAndServe() //if server is not able to start the it will returning the error otherwise it will returning the nil
		if err != nil {
			log.Fatal("Faild to Start Server !")
		}
	}()
	<-done

	slog.Info("Shutting Down the Server .. ")
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second) //we are provide the fime to shudown but also not shutting down
	defer cancel()
	err = server.Shutdown(cxt)
	if err != nil {
		slog.Error("Faild To Shutting down server", slog.String("Error :", err.Error()))
	}

	slog.Info("Serever Shutting down successfully !")
}
