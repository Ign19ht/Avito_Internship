package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateConnection() (*sqlx.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	source := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d", user, password, dbName, host, 5432)
	conn, err := sqlx.Connect("postgres", source)
	return conn, err
}

func CreateTables() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer db.Close()
	schemaBalances := `CREATE TABLE balances (
    id integer,
    balance real,
    PRIMARY KEY(id)
);`
	schemaReservation := `CREATE TABLE reservation (
    id_user integer,
    id_service integer,
    id_order integer,
    amount real,
    status text,
    date text,
    PRIMARY KEY(id_user, id_service, id_order)
);`
	schemaHistory := `CREATE TABLE history (
    id_user integer,
    date text,
    amount real,
    message text,
    PRIMARY KEY(id_user, date)
);`
	schemaReports := `CREATE TABLE reports (
    date text,
    link text,
    PRIMARY KEY(date)
);`
	db.Exec(schemaBalances)
	db.Exec(schemaReservation)
	db.Exec(schemaReports)
	db.Exec(schemaHistory)
}
