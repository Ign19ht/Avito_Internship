package service

import (
	"Avito/backend/repository"
	"Avito/backend/schemas"
	"net/http"
)

func Reservation(idUser int, idService int, idOrder int, amount float32, date string) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if _, err := repository.GetStatus(conn, idUser, idService, idOrder); err == nil {
		return http.StatusOK, schemas.ErrorResponse{Message: "Reservation already exist"}
	}

	balance, err := repository.GetBalance(conn, idUser)
	if err != nil {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "User not found"}
	}

	if balance >= amount {
		err = repository.AddReserve(conn, idUser, idService, idOrder, amount, date)
		if err == nil {
			return http.StatusOK, schemas.ErrorResponse{Message: "Money reserved"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}
	} else {
		return http.StatusBadRequest, schemas.ErrorResponse{Message: "Not enough money"}
	}
}

func ReservationConfirm(idUser int, idService int, idOrder int, amount float32, date string, message string) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	status, err := repository.GetStatus(conn, idUser, idService, idOrder)
	if err != nil {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}

	if status == "RESERVED" {
		amountInReserve, err := repository.GetAmountOfReserve(conn, idUser, idService, idOrder)
		if err != nil {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}
		if amountInReserve != amount {
			return http.StatusBadRequest, schemas.ErrorResponse{Message: "The amount of money doesn't match"}
		}
		err = repository.UpdateReserveStatus(conn, idUser, idService, idOrder,
			"CONFIRMED", date)
		if err == nil {
			repository.AddToHistory(conn, idUser, date, -amount, message)
			return http.StatusOK, schemas.ErrorResponse{Message: "Reservation confirmed"}
		} else {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}
	} else if status == "CONFIRMED" {
		return http.StatusOK, schemas.ErrorResponse{Message: "Reservation is already confirmed"}
	} else {
		return http.StatusBadRequest, schemas.ErrorResponse{Message: "Reservation is already canceled"}
	}
}

func ReservationCancel(idUser int, idService int, idOrder int, date string) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	status, err := repository.GetStatus(conn, idUser, idService, idOrder)
	if err != nil {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}

	if status == "RESERVED" {
		amount, err := repository.GetAmountOfReserve(conn, idUser, idService, idOrder)
		if err != nil {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		}

		err = repository.CancelReserve(conn, idUser, idService, idOrder, amount,
			date)
		if err != nil {
			return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
		} else {
			return http.StatusOK, schemas.ErrorResponse{Message: "Reserve canceled"}
		}
	} else {
		return http.StatusBadRequest, schemas.ErrorResponse{Message: "Reserve is already confirmed or canceled"}
	}
}
