package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/tejasthonge/Go-Backend/9_postgres_sql/1_setup/internals/config"
)

type Postgres struct {
	db *sql.DB
}

func NewConnectToDB(cfg *config.Config) *Postgres {
	// str:= "user=your_username password=your_password dbname=your_database sslmode=disable"
	var postgres Postgres
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.DbConfig.UserName, cfg.DbConfig.Password, cfg.DbConfig.DbName)
	
	sqlDb, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error to open datbase", err)
	}
	postgres = Postgres{
		db: sqlDb,
	}
	slog.Info("Database Created Succesfully!")
	return &postgres

}

func (p *Postgres) NewCreateTableUser() {

	creteTableQurry := `CREATE TABLE IF NOT EXIT(
		id SERIAL PRIMERY KEY
		name TEXT
		age INTEGER
	)`
	_, err := p.db.Exec(creteTableQurry)
	if err != nil {
		log.Fatal("Error To Create Users Table", err)
	}
	slog.Info("Users Table Created Successfully !")
}
