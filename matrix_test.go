package matrix_test

import (
	"testing"

	"github.com/EclesioMeloJunior/matrix"
	"github.com/stretchr/testify/assert"
)

func TestZeroMatrix(t *testing.T) {
	expectedM := matrix.Matrix(
		[][]int64{
			{0, 0},
			{0, 0},
		},
	)

	z := matrix.Zero(2, 2)

	assert.Equal(t, expectedM, *z)
}

func TestTransposeMatrix(t *testing.T) {
	expectedM := matrix.Matrix(
		[][]int64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	)

	m := matrix.Matrix(
		[][]int64{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	)

	m.Transpose()

	assert.Equal(t, expectedM, m)
}

func benchmarckTranspose(r, c int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := matrix.RandInt(r, c)
		m.Transpose()
	}
}

func BenchmarkTranspose3by3(b *testing.B) {
	benchmarckTranspose(3, 3, b)
}

func BenchmarkTranspose5by5(b *testing.B) {
	benchmarckTranspose(5, 5, b)
}

func BenchmarkTranspose10by10(b *testing.B) {
	benchmarckTranspose(10, 10, b)
}

func TestMatrixIsEqualsToAnother(t *testing.T) {
	a := matrix.RandInt(3, 3)
	b := a.Copy()

	assert.True(t, a.Equals(b))
}

func TestMatrixIsNotEqualsToAnother(t *testing.T) {
	a := matrix.RandInt(3, 3)
	b := a.Copy()
	*b = (*b)[:b.Rows()-1]

	assert.False(t, a.Equals(b))
}
