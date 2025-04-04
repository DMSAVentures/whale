package matrixoperations

import (
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

func (m Matrix) Invert() Matrix {
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

func (m Matrix) Flatten() string {
	var flat []string
	for _, row := range m {
		for _, val := range row {
			flat = append(flat, strconv.Itoa(val))
		}
	}

	return strings.Join(flat, ",")
}

func (m Matrix) Sum() int {
	var sum = 0
	for _, row := range m {
		for _, val := range row {
			sum += val
		}
	}

	return sum
}

func (m Matrix) Multiply() int {
	var product = 1
	for _, row := range m {
		for _, val := range row {
			product *= val
		}
	}

	return product
}
