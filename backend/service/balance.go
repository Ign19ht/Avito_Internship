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

	if _, err := db.GetBalance(conn, id); err != nil {
		err = db.CreateAccount(conn, id, amount)
		if err != err {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		} else {
			db.AddToHistory(conn, id, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money is credited"}
		}
	} else {
		err = db.AccrualMoney(conn, id, amount)
		if err == nil {
			db.AddToHistory(conn, id, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money is credited"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}
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

	balance, err := db.GetBalance(conn, idFrom)
	if err != nil {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}
	if balance < amount {
		return http.StatusBadRequest, schemas.ErrorResponse{Message: "Not enough money"}
	}

	if _, err := db.GetBalance(conn, idTo); err == nil {
		if err := db.SendMoney(conn, idFrom, idTo, amount); err == nil {
			db.AddToHistory(conn, idFrom, date, -amount, message)
			db.AddToHistory(conn, idTo, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
		}
	} else {
		if err := db.SendMoneyWithCreating(conn, idFrom, idTo, amount); err == nil {
			db.AddToHistory(conn, idFrom, date, -amount, message)
			db.AddToHistory(conn, idTo, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
		}
	}

	if err := db.SendMoney(conn, idFrom, idTo, amount); err == nil {
		db.AddToHistory(conn, idFrom, date, -amount, message)
		db.AddToHistory(conn, idTo, date, amount, message)
		return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
	} else {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
	}

}
