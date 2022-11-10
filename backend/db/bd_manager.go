package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateConnection() (*sqlx.DB, error) {
	//databaseUrl := "postgres://postgres:qwerlodaza@localhost:5432/postgres"
	conn, err := sqlx.Connect("postgres", "user=postgres password=qwerlodaza dbname=postgres sslmode=disable")
	return conn, err
}

func CreateTables() {
	db, err := CreateConnection()
	if err != nil {
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
