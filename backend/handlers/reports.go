package handlers

import (
	"Avito/backend/schemas"
	"Avito/backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetReport(c *gin.Context) {
	var request schemas.ReportRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	if request.Month > 12 || request.Month < 1 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.GetReport(request.Month, request.Year)
	c.JSON(code, response)
}

func GetHistory(c *gin.Context) {
	var request schemas.HistoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	if (request.SortType != "date" && request.SortType != "amount") || request.Amount < 1 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}
	if request.OrderType != "asc" && request.OrderType != "desc" {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Validation Failed"})
		return
	}

	code, response := service.GetHistory(request.Id, request.SortType, request.OrderType, request.Amount, request.Offset)
	c.JSON(code, response)
}
