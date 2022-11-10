package db

import (
	"Avito/backend/models"
	"github.com/jmoiron/sqlx"
)

func GetBalance(db *sqlx.DB, id int) (float32, error) {
	var balance models.Balance
	err := db.Get(&balance, "SELECT * from balances where id=$1", id)
	if err != nil {
		return 0, err
	}
	return balance.Balance, nil
}

func AccrualMoney(db *sqlx.DB, id int, amount float32) error {
	tx := db.MustBegin()
	tx.MustExec("update balances set balance=balance + $1 where id=$2", amount, id)
	err := tx.Commit()
	return err
}

func SendMoney(db *sqlx.DB, idFrom int, idTo int, amount float32) error {
	tx := db.MustBegin()
	tx.MustExec("update balances set balance=balance + $1 where id=$2", amount, idTo)
	tx.MustExec("update balances set balance=balance - $1 where id=$2", amount, idFrom)
	err := tx.Commit()
	return err
}
