package service

import (
	"Avito/backend/db"
	"Avito/backend/schemas"
	"net/http"
)

func MoneyAccrual(id int, amount float32, date string, message string) (int, interface{}) {
	conn, err := db.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if err := db.AccrualMoney(conn, id, amount); err == nil {
		db.AddToHistory(conn, id, date, amount, message)
		return http.StatusOK, schemas.ErrorResponse{Message: "Money is credited"}
	} else {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
}

func GetBalance(id int) (int, interface{}) {
	conn, err := db.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if balance, err := db.GetBalance(conn, id); err == nil {
		return http.StatusOK, schemas.BalanceResponse{Balance: balance}
	} else {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}
}

func SendMoney(idFrom int, idTo int, amount float32, date string, message string) (int, interface{}) {
	conn, err := db.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if err := db.SendMoney(conn, idFrom, idTo, amount); err == nil {
		db.AddToHistory(conn, idFrom, date, -amount, message)
		db.AddToHistory(conn, idTo, date, amount, message)
		return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
	} else {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}

}
