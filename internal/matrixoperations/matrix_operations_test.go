package matrixoperations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericMatrix_Invert(t *testing.T) {
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

	assert.Equal(t, len(expected), len(matrix), "Matrices have different number of rows")
	for i := range expected {
		assert.Equal(t, len(expected[i]), len(matrix[i]), "Matrices have different number of columns in row %d", i)
		for j := range expected[i] {
			assert.Equal(t, expected[i][j], matrix[i][j], "Matrices differ at element [%d][%d]", i, j)
		}
	}
}

func TestNumericMatrix_Flatten(t *testing.T) {
	// Test the flattenMatrix function
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := "1,2,3,4,5,6"
	flat := matrix.Flatten()

	assert.Equal(t, expected, flat)
}

func TestNumericMatrix_Sum(t *testing.T) {
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

func TestNumericMatrix_Multiply(t *testing.T) {
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

func TestNumericMatrix_String(t *testing.T) {
	// Test the String method
	matrix := NumericMatrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := "1,2,3\n4,5,6\n"
	result := matrix.String()

	assert.Equal(t, expected, result)
}

func TestAlphanumericMatrix_Invert(t *testing.T) {
	// Test the invertMatrix function
	matrix := AlphanumericMatrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := AlphanumericMatrix{
		{"1", "4", "7"},
		{"2", "5", "8"},
		{"3", "6", "9"},
	}
	matrix.Invert()

	assert.Equal(t, len(expected), len(matrix), "Matrices have different number of rows")
	for i := range expected {
		assert.Equal(t, len(expected[i]), len(matrix[i]), "Matrices have different number of columns in row %d", i)
		for j := range expected[i] {
			assert.Equal(t, expected[i][j], matrix[i][j], "Matrices differ at element [%d][%d]", i, j)
		}
	}
}

func TestAlphanumericMatrix_Flatten(t *testing.T) {
	// Test the flattenMatrix function
	matrix := AlphanumericMatrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	expected := "1,2,3,4,5,6"
	flat := matrix.Flatten()

	assert.Equal(t, expected, flat)
}

func TestAlphanumericMatrix_Sum(t *testing.T) {
	// Test the sumMatrix function
	matrix := AlphanumericMatrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	sum, err := matrix.Sum()

	assert.ErrorIs(t, err, ErrUnsupportedOperation)
	assert.Equal(t, 0, sum)
}

func TestAlphanumericMatrix_Multiply(t *testing.T) {
	// Test the multiplyMatrix function
	matrix := AlphanumericMatrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	product, err := matrix.Multiply()

	assert.ErrorIs(t, err, ErrUnsupportedOperation)
	assert.Equal(t, 0, product)
}

func TestAlphanumericMatrix_String(t *testing.T) {
	// Test the String method
	matrix := AlphanumericMatrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	expected := "1,2,3\n4,5,6\n"
	result := matrix.String()

	assert.Equal(t, expected, result)
}
