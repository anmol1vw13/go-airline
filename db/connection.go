package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5600
	user     = "airline_user"
	password = "airline_password"
	dbname   = "airline"
)

func Initiate() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(95)
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(time.Minute * 5)
    db.SetConnMaxIdleTime(time.Minute)
	return db;
}