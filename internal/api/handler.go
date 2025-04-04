package api

import (
	"encoding/csv"
	"fmt"
	"league/internal/utils"
	"net/http"
)

type MatrixProcessor interface {
	String() string
	Flatten() string
	Invert()
	Sum() (int64, error)
	Multiply() (int64, error)
}

// parseMatrix tries to parse [][]string as MatrixProcessor
func parseMatrix(data [][]string) (MatrixProcessor, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("matrix is empty")
	}

	// Try int parsing first
	intMatrix, err := utils.ParseIntMatrix(data)
	if err == nil {
		return &intMatrix, nil
	}

	// Fallback to string
	stringMatrix, err := utils.ParseStringMatrix(data)
	if err == nil {
		return &stringMatrix, nil
	}

	return nil, fmt.Errorf("unable to parse matrix as int type or string type")
}

func parseCSVFromRequest(r *http.Request) ([][]string, error) {
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
	records, err := parseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := parseMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(w, 200, matrix.String())
}

func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := parseMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	matrix.Invert()

	respond(w, 200, matrix.String())
}

func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := parseMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flat := matrix.Flatten()
	respond(w, 200, flat)
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := parseMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sum, err := matrix.Sum()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respond(w, 200, sum)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := parseCSVFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := parseMatrix(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := matrix.Multiply()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to process request: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	respond(w, 200, product)
}

func respond(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	if _, err := fmt.Fprintln(w, body); err != nil {
		// log error
		fmt.Printf("failed to write response: %v\n", err)
	}
}
