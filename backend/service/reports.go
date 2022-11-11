package service

import (
	"Avito/backend/fileWriter"
	"Avito/backend/repository"
	"Avito/backend/schemas"
	"fmt"
	"net/http"
	"strconv"
)

func GetReport(month int, year int) (int, interface{}) {
	monthStr := strconv.Itoa(month)
	yearStr := strconv.Itoa(year)
	date := fmt.Sprintf("%d-%d", year, month)

	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	if link, err := repository.GetReport(conn, date); err == nil {
		return http.StatusOK, schemas.ReportResponse{Link: link}
	}

	data, err := repository.GetMonthlyReport(conn, monthStr, yearStr)
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	if len(data) == 0 {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}

	report := map[int]float32{}
	for _, item := range data {
		if _, ok := report[item.IdService]; ok {
			report[item.IdService] += item.Amount
		} else {
			report[item.IdService] = item.Amount
		}
	}

	if link, err := fileWriter.WriteCSV(date, report); err == nil {
		repository.AddReport(conn, date, link)
		return http.StatusOK, schemas.ReportResponse{Link: link}
	} else {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "Write file error"}
	}

}

func GetHistory(id int, sortType string, orderType string, amount int, offset int) (int, interface{}) {
	conn, err := repository.CreateConnection()
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	defer conn.Close()

	history, err := repository.GetHistory(conn, id, sortType, orderType, amount, offset)
	if err != nil {
		return http.StatusInternalServerError, schemas.ErrorResponse{Message: "DataBase error"}
	}
	if len(history) == 0 {
		return http.StatusNotFound, schemas.ErrorResponse{Message: "Item not found"}
	}

	var response []schemas.HistoryResponseItem
	for _, item := range history {
		response = append(response,
			schemas.HistoryResponseItem{Date: item.Date, Amount: item.Amount, Message: item.Message})
	}
	return http.StatusOK, schemas.HistoryResponse{History: response}
}
