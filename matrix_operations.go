package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) String() string {
	var output string
	for _, row := range m {
		strRow := make([]string, len(row))
		for j, val := range row {
			strRow[j] = strconv.Itoa(val)
		}

		output = output + strings.Join(strRow, ",") + "\n"
	}
	return output
}

func (m Matrix) invertMatrix() Matrix {
	size := len(m)
	inverted := make(Matrix, size)
	for i := 0; i < size; i++ {
		inverted[i] = make([]int, size)
		for j := 0; j < size; j++ {
			inverted[i][j] = m[j][i]
		}
	}
	return inverted
}

func (m Matrix) flattenMatrix() string {
	var flat []string
	for _, row := range m {
		for _, val := range row {
			flat = append(flat, strconv.Itoa(val))
		}
	}
	return strings.Join(flat, ",")
}

func (m Matrix) sumMatrix() int {
	var sum = 0
	for _, row := range m {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

func (m Matrix) multiplyMatrix() int {
	var product = 1
	for _, row := range m {
		for _, val := range row {
			product *= val
		}
	}
	return product
}

func ParseToMatrix(data [][]string) (Matrix, error) {
	var matrix Matrix
	for _, row := range data {
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
