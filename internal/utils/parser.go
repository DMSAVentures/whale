package utils

import (
	"fmt"
	"league/internal/matrixoperations"
	"strconv"
)

func ParseToMatrix(data [][]string) (matrixoperations.Matrix, error) {
	var matrix matrixoperations.Matrix

	rowLength := len(data[0])

	for _, row := range data {
		// Check if all rows have the same length
		if len(row) != rowLength {
			return nil, fmt.Errorf("failed to parse csv records to matrix: all rows must have the same length")
		}

		// Parse each row to integers
		var intRow []int
		for _, val := range row {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("failed to convert string to int: %w", err)
			}
			intRow = append(intRow, intVal)
		}

		matrix = append(matrix, intRow)
	}

	return matrix, nil
}
