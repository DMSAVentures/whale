package matrixoperations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertMatrixEqual(t *testing.T, expected, actual NumericMatrix) {
	assert.Equal(t, len(expected), len(actual), "Matrices have different number of rows")
	for i := range expected {
		assert.Equal(t, len(expected[i]), len(actual[i]), "Matrices have different number of columns in row %d", i)
		for j := range expected[i] {
			assert.Equal(t, expected[i][j], actual[i][j], "Matrices differ at element [%d][%d]", i, j)
		}
	}
}

func TestInvertMatrix(t *testing.T) {
	// Test the invertMatrix function
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := NumericMatrix{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	matrix.Invert()

	assertMatrixEqual(t, expected, matrix)
}

func TestFlattenMatrix(t *testing.T) {
	// Test the flattenMatrix function
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := "1,2,3,4,5,6"
	flat := matrix.Flatten()

	assert.Equal(t, expected, flat)
}

func TestSumMatrix(t *testing.T) {
	// Test the sumMatrix function
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := 21
	sum, err := matrix.Sum()

	assert.NoError(t, err)
	assert.Equal(t, expected, sum)
}

func TestMultiplyMatrix(t *testing.T) {
	// Test the multiplyMatrix function
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := 720
	product, err := matrix.Multiply()

	assert.NoError(t, err)
	assert.Equal(t, expected, product)
}

func TestMatrixToString(t *testing.T) {
	// Test the String method
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := "1,2,3\n4,5,6\n"
	result := matrix.String()

	assert.Equal(t, expected, result)
}
