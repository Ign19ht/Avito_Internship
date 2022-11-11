package repository

import (
	"Avito/backend/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func AddToHistory(db *sqlx.DB, idUser int, date string, amount float32, message string) error {
	tx := db.MustBegin()
	tx.MustExec("insert into history (id_user, date, amount, message) values ($1, $2, $3, $4)",
		idUser, date, amount, message)
	err := tx.Commit()
	return err
}

func GetHistory(db *sqlx.DB, idUser int, sortType string, orderType string, amount int, offset int) ([]models.History, error) {
	var history []models.History
	query := fmt.Sprintf("SELECT * from history where id_user=%d order by %s %s limit %d offset %d",
		idUser, sortType, orderType, amount, offset)
	err := db.Select(&history, query)
	if err != nil {
		return nil, err
	}
	return history, nil
}
