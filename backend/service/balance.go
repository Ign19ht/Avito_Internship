package service

import (
	"Avito/backend/repository"
	"Avito/backend/schemas"
	"net/http"
)

func MoneyAccrual(id int, amount float32, date string, message string) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if _, err := repository.GetBalance(conn, id); err != nil {
		err = repository.CreateAccount(conn, id, amount)
		if err != err {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		} else {
			repository.AddToHistory(conn, id, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money is credited"}
		}
	} else {
		err = repository.AccrualMoney(conn, id, amount)
		if err == nil {
			repository.AddToHistory(conn, id, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money is credited"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}
	}
}

func GetBalance(id int) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if balance, err := repository.GetBalance(conn, id); err == nil {
		return http.StatusOK, schemas.BalanceResponse{Balance: balance}
	} else {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}
}

func SendMoney(idFrom int, idTo int, amount float32, date string, message string) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	balance, err := repository.GetBalance(conn, idFrom)
	if err != nil {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}
	if balance < amount {
		return http.StatusBadRequest, schemas.ErrorResponse{Message: "Not enough money"}
	}

	if _, err := repository.GetBalance(conn, idTo); err == nil {
		if err := repository.SendMoney(conn, idFrom, idTo, amount); err == nil {
			repository.AddToHistory(conn, idFrom, date, -amount, message)
			repository.AddToHistory(conn, idTo, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
		}
	} else {
		if err := repository.SendMoneyWithCreating(conn, idFrom, idTo, amount); err == nil {
			repository.AddToHistory(conn, idFrom, date, -amount, message)
			repository.AddToHistory(conn, idTo, date, amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
		}
	}

	if err := repository.SendMoney(conn, idFrom, idTo, amount); err == nil {
		repository.AddToHistory(conn, idFrom, date, -amount, message)
		repository.AddToHistory(conn, idTo, date, amount, message)
		return http.StatusOK, schemas.ErrorResponse{Message: "Money sent"}
	} else {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Database error"}
	}

}
