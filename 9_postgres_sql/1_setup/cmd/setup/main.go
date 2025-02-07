package main

import (
	"github.com/tejasthonge/Go-Backend/9_postgres_sql/1_setup/internals/config"
	"github.com/tejasthonge/Go-Backend/9_postgres_sql/1_setup/internals/database/postgres"
)

func main() {

	dbCfg := config.DbConfig{
		UserName: "boss",
		Password: "password",
		DbName:   "setup",
	}
	cfg := config.Config{
		Env:      "Dev",
		Addr:     "8055",
		DbConfig: dbCfg,
	}

	psqlStorage := postgres.NewConnectToDB(&cfg) //here we coonectiong the database
	psqlStorage.NewCreateTableUser()



}
