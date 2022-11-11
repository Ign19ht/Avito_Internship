package repository

import (
	"Avito/backend/models"
	"github.com/jmoiron/sqlx"
)

func AddReport(db *sqlx.DB, date string, link string) error {
	tx := db.MustBegin()
	tx.MustExec("insert into reports (date, link) values ($1, $2)",
		date, link)
	err := tx.Commit()
	return err
}

func GetReport(db *sqlx.DB, date string) (string, error) {
	var history models.Reports
	err := db.Get(&history, "SELECT * from reports where date=$1",
		date)
	if err != nil {
		return "", err
	}
	return history.Link, nil
}
