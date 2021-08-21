package matrix_test

import (
	"testing"

	"github.com/EclesioMeloJunior/matrix"
	"github.com/stretchr/testify/assert"
)

func TestDeterminantsOfManyMatrix(t *testing.T) {
	testCases := []struct {
		m         matrix.Matrix
		exepected int64
	}{
		{
			m: matrix.Matrix([][]int64{
				{8, 14},
				{6, 13},
			}),
			exepected: 20,
		},
		{
			m: matrix.Matrix([][]int64{
				{2, 1, 2},
				{3, 4, 5},
				{6, 7, 4},
			}),
			exepected: -26,
		},
		{
			m: matrix.Matrix([][]int64{
				{2, 1, 2, 4},
				{3, 4, 1, 1},
				{6, 7, 4, 2},
				{1, 1, 2, 4},
			}),
			exepected: 32,
		},
	}

	for _, tcase := range testCases {
		v, err := tcase.m.Determinant()
		assert.NoError(t, err)
		assert.Equal(t, tcase.exepected, v)
	}
}

func TestGetDeterminant3x3Subset(t *testing.T) {
	m := matrix.Matrix([][]int64{
		{1, 2, 4},
		{3, 8, 14},
		{2, 6, 13},
	})

	s := matrix.GetDeterminantSubset(&m, 0)
	expected := matrix.Matrix([][]int64{
		{8, 14},
		{6, 13},
	})
	assert.Equal(t, &expected, s)

	s = matrix.GetDeterminantSubset(&m, 1)
	expected = matrix.Matrix([][]int64{
		{3, 14},
		{2, 13},
	})
	assert.Equal(t, &expected, s)

	s = matrix.GetDeterminantSubset(&m, 2)
	expected = matrix.Matrix([][]int64{
		{3, 8},
		{2, 6},
	})
	assert.Equal(t, &expected, s)
}

func TestGetDeterminant4x4Subset(t *testing.T) {
	m := matrix.Matrix([][]int64{
		{1, 2, 4, 5},
		{3, 8, 14, 6},
		{2, 6, 13, 0},
		{2, 90, 53, 10},
	})

	testCases := []struct {
		column   int
		expected matrix.Matrix
	}{
		{
			column: 0,
			expected: matrix.Matrix([][]int64{
				{8, 14, 6},
				{6, 13, 0},
				{90, 53, 10},
			}),
		},
		{
			column: 1,
			expected: matrix.Matrix([][]int64{
				{3, 14, 6},
				{2, 13, 0},
				{2, 53, 10},
			}),
		},
		{
			column: 2,
			expected: matrix.Matrix([][]int64{
				{3, 8, 6},
				{2, 6, 0},
				{2, 90, 10},
			}),
		},
		{
			column: 3,
			expected: matrix.Matrix([][]int64{
				{3, 8, 14},
				{2, 6, 13},
				{2, 90, 53},
			}),
		},
	}

	for _, tcase := range testCases {
		s := matrix.GetDeterminantSubset(&m, tcase.column)
		assert.Equal(t, &tcase.expected, s)
	}
}
