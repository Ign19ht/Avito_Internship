package handlers

import (
	"Avito/backend/schemas"
	"Avito/backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Reservation(c *gin.Context) {
	var request schemas.ReservationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	_, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.Reservation(request.IdUser, request.IdService, request.IdOrder, request.Amount, request.Date)
	c.JSON(code, response)
}

func ReservationConfirm(c *gin.Context) {
	var request schemas.ConfirmRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	_, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.ReservationConfirm(request.IdUser, request.IdService,
		request.IdOrder, request.Amount, request.Date, request.Message)
	c.JSON(code, response)
}

func ReservationCancel(c *gin.Context) {
	var request schemas.CancelRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	_, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.ReservationCancel(request.IdUser, request.IdService, request.IdOrder, request.Date)
	c.JSON(code, response)
}
