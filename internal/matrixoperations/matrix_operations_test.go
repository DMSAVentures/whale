package matrixoperations

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock large matrix (100x100)
var largeNumericMatrix = func() NumericMatrix {
	matrix := make(NumericMatrix, 100)
	for i := range matrix {
		matrix[i] = make([]int, 100)
		for j := range matrix[i] {
			matrix[i][j] = (i*j + 1) % 9 // ensure small numbers to prevent overflow
			if matrix[i][j] == 0 {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}()

func TestNumericMatrix_Invert(t *testing.T) {
	tests := []struct {
		name     string
		matrix   NumericMatrix
		expected NumericMatrix
	}{
		{
			"3x3 matrix",
			NumericMatrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			NumericMatrix{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			"2x2 matrix",
			NumericMatrix{{1, 2}, {3, 4}},
			NumericMatrix{{1, 3}, {2, 4}},
		},
		{
			"1x1 matrix",
			NumericMatrix{{42}},
			NumericMatrix{{42}},
		},
		{
			"2x3 matrix",
			NumericMatrix{{1, 2, 3}, {4, 5, 6}},
			NumericMatrix{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			"Empty matrix",
			NumericMatrix{},
			NumericMatrix{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.matrix.Invert()
			assert.Equal(t, tt.expected, tt.matrix)
		})
	}
}

func TestNumericMatrix_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		matrix   NumericMatrix
		expected string
	}{
		{"Empty matrix", NumericMatrix{}, ""},
		{"1x1 matrix", NumericMatrix{{9}}, "9"},
		{"2x2 matrix", NumericMatrix{{1, 2}, {3, 4}}, "1,2,3,4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.matrix.Flatten())
		})
	}
}

func TestNumericMatrix_Multiply(t *testing.T) {
	tests := []struct {
		name     string
		matrix   NumericMatrix
		expected int64
		wantErr  bool
	}{
		{"1x1 matrix", NumericMatrix{{7}}, 7, false},
		{"2x2 matrix", NumericMatrix{{2, 3}, {4, 5}}, 120, false},
		{"Edge values", NumericMatrix{{1, 1}, {1, 1}}, 1, false},
		{"3x3 matrix", NumericMatrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 362880, false},
		{"Partial large matrix", NumericMatrix{{2, 3, 4, 5, 6}, {1, 1, 1, 1, 1}}, 720, false},
		{
			"Mixed negatives",
			NumericMatrix{{-1, 2}, {-3, 4}},
			24, false, // -1 * 2 * -3 * 4 = 24
		},
		{
			"Overflow multiply",
			NumericMatrix{{math.MaxInt64, 2}},
			0, true,
		},
		{
			"Standard 7x3 matrix",
			NumericMatrix{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
				{10, 11, 12}, {13, 14, 15},
				{16, 17, 18}, {19, 20, 21},
			},
			0, true, // Product would exceed int64
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := tt.matrix.Multiply()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, product)
			}
		})
	}
}

func TestNumericMatrix_Sum(t *testing.T) {
	tests := []struct {
		name     string
		matrix   NumericMatrix
		expected int64
		wantErr  bool
	}{
		{"Empty matrix", NumericMatrix{}, 0, false},
		{"1x1 matrix", NumericMatrix{{42}}, 42, false},
		{"Negative values", NumericMatrix{{-1, -2}, {3, 4}}, 4, false},
		{"2x3 matrix", NumericMatrix{{1, 2, 3}, {4, 5, 6}}, 21, false},
		{"Large 100x100 matrix", largeNumericMatrix, 36862, false},
		{
			"Mixed negative and positive",
			NumericMatrix{{1, -2, 3}, {-4, 5, -6}, {7, -8, 9}},
			5, false,
		},
		{
			"Overflow sum",
			NumericMatrix{{math.MaxInt64, 1}},
			0, true,
		},
		{
			"Standard 7x3 matrix",
			NumericMatrix{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
				{10, 11, 12}, {13, 14, 15},
				{16, 17, 18}, {19, 20, 21},
			},
			231, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, err := tt.matrix.Sum()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, sum)
			}
		})
	}
}

func TestNumericMatrix_String(t *testing.T) {
	matrix := NumericMatrix{{1, 2, 3}, {4, 5, 6}}
	expected := "1,2,3\n4,5,6\n"
	result := matrix.String()
	assert.Equal(t, expected, result)
}

func TestAlphanumericMatrix_Invert(t *testing.T) {
	tests := []struct {
		name     string
		matrix   AlphanumericMatrix
		expected AlphanumericMatrix
	}{
		{
			name: "3x3 matrix",
			matrix: AlphanumericMatrix{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expected: AlphanumericMatrix{
				{"1", "4", "7"},
				{"2", "5", "8"},
				{"3", "6", "9"},
			},
		},
		{
			name:     "Empty matrix",
			matrix:   AlphanumericMatrix{},
			expected: AlphanumericMatrix{},
		},
		{
			name: "1x1 matrix",
			matrix: AlphanumericMatrix{
				{"X"},
			},
			expected: AlphanumericMatrix{
				{"X"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.matrix.Invert()
			assert.Equal(t, tt.expected, tt.matrix)
		})
	}
}

func TestAlphanumericMatrix_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		matrix   AlphanumericMatrix
		expected string
	}{
		{"Empty matrix", AlphanumericMatrix{}, ""},
		{"1x1 matrix", AlphanumericMatrix{{"A"}}, "A"},
		{"2x2 matrix", AlphanumericMatrix{{"1", "2"}, {"3", "4"}}, "1,2,3,4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.matrix.Flatten())
		})
	}
}

func TestAlphanumericMatrix_Sum(t *testing.T) {
	matrix := AlphanumericMatrix{{"1", "2", "3"}, {"4", "5", "6"}}
	sum, err := matrix.Sum()
	assert.ErrorIs(t, err, ErrUnsupportedOperation)
	assert.Equal(t, int64(0), sum)
}

func TestAlphanumericMatrix_Multiply(t *testing.T) {
	matrix := AlphanumericMatrix{{"1", "2", "3"}, {"4", "5", "6"}}
	product, err := matrix.Multiply()
	assert.ErrorIs(t, err, ErrUnsupportedOperation)
	assert.Equal(t, int64(0), product)
}

func TestAlphanumericMatrix_String(t *testing.T) {
	matrix := AlphanumericMatrix{{"1", "2", "3"}, {"4", "5", "6"}}
	expected := "1,2,3\n4,5,6\n"
	result := matrix.String()
	assert.Equal(t, expected, result)
}
