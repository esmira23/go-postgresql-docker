package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

const (
	host     = "postgresdb"
	port     = "5432"
	user     = "esmira"
	password = "secret"
	dbname   = "esmira"
	sslmode  = "disable"
)

func ConnectDB() {
	var err error

	psqconn := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	DB, err = sqlx.Open("postgres", psqconn)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

}
