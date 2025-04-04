package matrixoperations

import (
	"errors"
	"strconv"
	"strings"
)

var ErrUnsupportedOperation = errors.New("unsupported operation")

type NumericMatrix [][]int

func (m *NumericMatrix) String() string {
	var output string
	for _, row := range *m {
		strRow := make([]string, len(row))
		for j, val := range row {
			strRow[j] = strconv.Itoa(val)
		}

		output = output + strings.Join(strRow, ",") + "\n"
	}

	return output
}

func (m *NumericMatrix) Invert() {
	size := len(*m)
	inverted := make(NumericMatrix, size)
	for i := 0; i < size; i++ {
		inverted[i] = make([]int, size)
		for j := 0; j < size; j++ {
			inverted[i][j] = (*m)[j][i]
		}
	}

	*m = inverted
}

func (m *NumericMatrix) Flatten() string {
	var flat []string
	for _, row := range *m {
		for _, val := range row {
			flat = append(flat, strconv.Itoa(val))
		}
	}

	return strings.Join(flat, ",")
}

func (m *NumericMatrix) Sum() (int, error) {
	var sum = 0
	for _, row := range *m {
		for _, val := range row {
			sum += val
		}
	}

	return sum, nil
}

func (m *NumericMatrix) Multiply() (int, error) {
	var product = 1
	for _, row := range *m {
		for _, val := range row {
			product *= val
		}
	}

	return product, nil
}

type AlphanumericMatrix [][]string

func (a *AlphanumericMatrix) String() string {
	var output string
	for _, row := range *a {
		strRow := make([]string, len(row))
		for j, val := range row {
			strRow[j] = val
		}

		output = output + strings.Join(strRow, ",") + "\n"
	}

	return output
}

func (a *AlphanumericMatrix) Flatten() string {
	var flat []string
	for _, row := range *a {
		for _, val := range row {
			flat = append(flat, val)
		}
	}

	return strings.Join(flat, ",")
}

func (a *AlphanumericMatrix) Invert() {
	size := len(*a)
	inverted := make(AlphanumericMatrix, size)
	for i := 0; i < size; i++ {
		inverted[i] = make([]string, size)
		for j := 0; j < size; j++ {
			inverted[i][j] = (*a)[j][i]
		}
	}

	*a = inverted
}

func (a *AlphanumericMatrix) Sum() (int, error) {
	return 0, ErrUnsupportedOperation
}

func (a *AlphanumericMatrix) Multiply() (int, error) {
	return 0, ErrUnsupportedOperation
}
