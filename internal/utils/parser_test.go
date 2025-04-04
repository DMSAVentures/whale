package utils

import (
	"league/internal/matrixoperations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntMatrix(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]string
		expected  matrixoperations.NumericMatrix
		expectErr bool
	}{
		{
			name:     "Valid numeric matrix",
			input:    [][]string{{"1", "2"}, {"3", "4"}},
			expected: matrixoperations.NumericMatrix{{1, 2}, {3, 4}},
		},
		{
			name:      "Inconsistent row length",
			input:     [][]string{{"1", "2"}, {"3"}},
			expectErr: true,
		},
		{
			name:      "Invalid integer",
			input:     [][]string{{"1", "a"}, {"3", "4"}},
			expectErr: true,
		},
		{
			name:     "Empty matrix",
			input:    [][]string{},
			expected: matrixoperations.NumericMatrix{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseIntMatrix(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestParseStringMatrix(t *testing.T) {
	tests := []struct {
		name      string
		input     [][]string
		expected  matrixoperations.AlphanumericMatrix
		expectErr bool
	}{
		{
			name:     "Valid string matrix",
			input:    [][]string{{"a", "b"}, {"c", "d"}},
			expected: matrixoperations.AlphanumericMatrix{{"a", "b"}, {"c", "d"}},
		},
		{
			name:      "Inconsistent row length",
			input:     [][]string{{"a", "b"}, {"c"}},
			expectErr: true,
		},
		{
			name:     "Empty matrix",
			input:    [][]string{},
			expected: matrixoperations.AlphanumericMatrix{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseStringMatrix(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
