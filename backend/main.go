package main

import (
	"Avito/backend/handlers"
	"Avito/backend/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.CreateTables()
	router := gin.Default()
	router.POST("/balance/accrual", handlers.MoneyAccrual)
	router.POST("/balance/send", handlers.SendMoney)
	router.GET("/balance/balance", handlers.GetBalance)
	router.POST("/order/reserve", handlers.Reservation)
	router.POST("/order/cancel", handlers.ReservationCancel)
	router.POST("/order/confirm", handlers.ReservationConfirm)
	router.GET("/report", handlers.GetReport)
	router.GET("/history", handlers.GetHistory)

	router.Run("0.0.0.0:8080")
}
