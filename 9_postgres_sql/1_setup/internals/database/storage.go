package database

import "github.com/tejasthonge/Go-Backend/9_postgres_sql/1_setup/internals/config"

type Storage interface{
	New(cfg *config.Config)
}