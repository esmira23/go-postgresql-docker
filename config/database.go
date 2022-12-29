package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "secret"
	dbname   = "golang-postgresql-docker"
	sslmode  = "disable"
)

func ConnectDB() {

	psqconn := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sqlx.Open("postgres", psqconn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	DB = db

}
