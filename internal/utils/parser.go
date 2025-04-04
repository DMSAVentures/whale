package utils

import (
	"fmt"
	"league/internal/matrixoperations"
	"strconv"
)

func ParseIntMatrix(data [][]string) (matrixoperations.NumericMatrix, error) {
	rowLen := len(data[0])
	matrix := make(matrixoperations.NumericMatrix, len(data))

	for i, row := range data {
		if len(row) != rowLen {
			return nil, fmt.Errorf("row %d has inconsistent length", i+1)
		}
		intRow := make([]int, rowLen)
		for j, val := range row {
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid int at row %d col %d: %w", i+1, j+1, err)
			}
			intRow[j] = n
		}
		matrix[i] = intRow
	}
	return matrix, nil
}

func ParseStringMatrix(data [][]string) (matrixoperations.AlphanumericMatrix, error) {
	rowLen := len(data[0])
	matrix := make(matrixoperations.AlphanumericMatrix, len(data))

	for i, row := range data {
		if len(row) != rowLen {
			return nil, fmt.Errorf("row %d has inconsistent length", i+1)
		}
		matrix[i] = append([]string(nil), row...)
	}
	return matrix, nil
}
