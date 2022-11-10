package handlers

import (
	"Avito/backend/schemas"
	"Avito/backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MoneyAccrual(c *gin.Context) {
	var request schemas.AccrualRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	_, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.MoneyAccrual(request.Id, request.Amount, request.Date, request.Message)
	c.JSON(code, response)

}

func GetBalance(c *gin.Context) {
	var request schemas.BalanceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.GetBalance(request.Id)
	c.JSON(code, response)
}

func SendMoney(c *gin.Context) {
	var request schemas.SendRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	_, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.SendMoney(request.IdFrom, request.IdTo, request.Amount, request.Date, request.Message)
	c.JSON(code, response)

}
