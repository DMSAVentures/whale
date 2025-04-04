package matrixoperations

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var ErrUnsupportedOperation = errors.New("unsupported operation")
var ErrOverflow = errors.New("integer overflow encountered")

type NumericMatrix [][]int

type AlphanumericMatrix [][]string

func safeMultiply(a, b int64) (int64, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	if a == math.MinInt64 || b == math.MinInt64 {
		// edge case where abs(MinInt64) overflows
		return 0, ErrOverflow
	}

	if absInt64(a) > math.MaxInt64/absInt64(b) {
		return 0, ErrOverflow
	}

	return a * b, nil
}

func absInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func safeAdd(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) ||
		(b < 0 && a < math.MinInt64-b) {
		return 0, ErrOverflow
	}
	return a + b, nil
}

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
	if size == 0 {
		return
	}
	rowLen := len((*m)[0])
	inverted := make(NumericMatrix, rowLen)
	for i := 0; i < rowLen; i++ {
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

func (m *NumericMatrix) Sum() (int64, error) {
	var sum int64 = 0
	for _, row := range *m {
		for _, val := range row {
			x, err := safeAdd(int64(sum), int64(val))
			if err != nil {
				return 0, err
			}
			sum = x
		}
	}

	return sum, nil
}

func (m *NumericMatrix) Multiply() (int64, error) {
	var product int64 = 1
	for _, row := range *m {
		for _, val := range row {
			x, err := safeMultiply(product, int64(val))
			if err != nil {
				return 0, err
			}
			product = x
		}
	}

	return product, nil
}

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
	if size == 0 {
		return
	}

	inverted := make(AlphanumericMatrix, size)
	for i := 0; i < size; i++ {
		inverted[i] = make([]string, size)
		for j := 0; j < size; j++ {
			inverted[i][j] = (*a)[j][i]
		}
	}

	*a = inverted
}

func (a *AlphanumericMatrix) Sum() (int64, error) {
	return 0, ErrUnsupportedOperation
}

func (a *AlphanumericMatrix) Multiply() (int64, error) {
	return 0, ErrUnsupportedOperation
}
