package main

import (
	"Avito/backend/db"
	"Avito/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.CreateTables()
	router := gin.Default()
	router.POST("/accrual", handlers.MoneyAccrual)
	router.POST("/reserve", handlers.Reservation)
	router.POST("/cancel", handlers.ReservationCancel)
	router.POST("/confirm", handlers.ReservationConfirm)
	router.POST("/send", handlers.SendMoney)
	router.GET("/balance", handlers.GetBalance)
	router.GET("/report", handlers.GetReport)
	router.GET("/history", handlers.GetHistory)

	router.Run("0.0.0.0:8080")
}
