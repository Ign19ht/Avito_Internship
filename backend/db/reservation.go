package db

import (
	"Avito/backend/models"
	"github.com/jmoiron/sqlx"
)

func AddReserve(db *sqlx.DB, idUser int, idService int, idOrder int, amount float32, date string) error {
	tx := db.MustBegin()
	tx.MustExec("insert into reservation (id_user, id_service, id_order, amount, status, date) values ($1, $2, $3, $4, $5, $6)",
		idUser, idService, idOrder, amount, "RESERVED", date)
	tx.MustExec("update balances set balance=balance - $1 where id=$2", amount, idUser)
	err := tx.Commit()
	return err
}

func CancelReserve(db *sqlx.DB, idUser int, idService int, idOrder int, amount float32, date string) error {
	tx := db.MustBegin()
	tx.MustExec("update reservation set status=$4, date=$5 where id_user=$1 and id_service=$2 and id_order=$3",
		idUser, idService, idOrder, "CANCEL", date)
	tx.MustExec("update balances set balance=balance + $1 where id=$2", amount, idUser)
	err := tx.Commit()
	return err
}

func UpdateReserveStatus(db *sqlx.DB, idUser int, idService int, idOrder int, newStatus string, date string) error {
	tx := db.MustBegin()
	tx.MustExec("update reservation set status=$4, date=$5 where id_user=$1 and id_service=$2 and id_order=$3",
		idUser, idService, idOrder, newStatus, date)
	err := tx.Commit()
	return err
}

func GetStatus(db *sqlx.DB, idUser int, idService int, idOrder int) (string, error) {
	var reservation models.Reservation
	err := db.Get(&reservation, "SELECT * from reservation where id_user=$1 and id_service=$2 and id_order=$3",
		idUser, idService, idOrder)
	if err != nil {
		return "", err
	}
	return reservation.Status, nil
}

func GetAmountOfReserve(db *sqlx.DB, idUser int, idService int, idOrder int) (float32, error) {
	var reservation models.Reservation
	err := db.Get(&reservation, "SELECT * from reservation where id_user=$1 and id_service=$2 and id_order=$3",
		idUser, idService, idOrder)
	if err != nil {
		return 0, err
	}
	return reservation.Amount, nil
}

func GetMonthlyReport(db *sqlx.DB, month string, year string) ([]models.Reservation, error) {
	var reservations []models.Reservation
	date := year + "-" + month + "-__T__:__:__Z"
	err := db.Select(&reservations, "SELECT * from reservation where date like $1 and status=$2", date, "CONFIRMED")
	if err != nil {
		return nil, err
	}

	return reservations, nil
}
