package fileWriter

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WriteCSV(date string, report map[int]float32) (string, error) {
	data := [][]string{{"id_service", "revenue"}}

	for key, value := range report {
		data = append(data, []string{strconv.Itoa(key), fmt.Sprintf("%.2f", value)})
	}

	link := fmt.Sprintf("%s/%s.csv", "reports", date)

	f, err := os.Create(link)
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
		return "", err
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(data)
	if err != nil {
		return "", err
	}

	return link, nil
}
