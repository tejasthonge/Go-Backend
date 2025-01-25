package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // we geting the error we importiing them but not use so we use _ ,the resion is we not use the driver directly it used behind the seen so we must have to add _ or we not implort them
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/config"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES (?,?,?)")//here we first 
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age)//here we passing the value to the statmet
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(cfg *config.Config) (*Sqlite, error) {
	//here we opening the we have pass two  things that is Driver , storage path in our case we user sqlite so this driwer comes form  ->go get github.com/mattn/go-sqlite3
	//if we use prostgress then we have to pass the postgress driver which are present in athor package we jusct have the serch on it postgress driver
	//tand this is returning error or the Db
	sqliteDb, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		log.Fatal("errer to open the Sql Database")
		return nil, err
	}

	//if not getting error we have to create table in db
	_, err = sqliteDb.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)
	if err != nil {
		return nil, err
	}
	//if the result is getting

	//now we have to show all the dabase in graficale use interfase
	//by useing of  tabable pluse
	return &Sqlite{
		Db: sqliteDb, //here we have to pass the adree but open method retun the adrees of db
	}, nil
}
