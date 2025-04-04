package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func ParseCSVFromRequest(r *http.Request) ([][]string, error) {
	var records [][]string
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to get file from request: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	return records, nil
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ParseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := ParseToMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, matrix.String())
}
