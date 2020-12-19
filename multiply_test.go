package matrix_test

import (
	"testing"

	"github.com/EclesioMeloJunior/matrix"
	"github.com/stretchr/testify/assert"
)

func TestMulltiplyByScalar(t *testing.T) {
	a := matrix.Matrix(
		[][]int64{
			{3, 3},
			{2, 2},
		},
	)
	multiplyBy := 3

	expected := matrix.Matrix(
		[][]int64{
			{9, 9},
			{6, 6},
		},
	)

	a.MultplyByScalar(int64(multiplyBy))

	assert.Equal(t, expected, a)
}

func TestMultiplyByMatrix(t *testing.T) {
	a := matrix.Matrix(
		[][]int64{
			{2, 5, 9},
			{3, 6, 8},
		},
	)

	b := matrix.Matrix(
		[][]int64{
			{2, 7},
			{4, 3},
			{5, 2},
		},
	)

	expected := matrix.Matrix(
		[][]int64{
			{69, 47},
			{70, 55},
		},
	)

	m, err := a.Multiply(&b)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func benchmarkMultiplyMatrix(ra, ca, rb, cb int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := matrix.RandInt(ra, ca)
		b := matrix.RandInt(rb, cb)

		a.Multiply(b)
	}
}

func BenchmarkMultiply3by5Times4by3(b *testing.B) {
	benchmarkMultiplyMatrix(3, 5, 4, 3, b)
}

func BenchmarkMultiply10by20Times20by60(b *testing.B) {
	benchmarkMultiplyMatrix(10, 20, 20, 60, b)
}

func BenchmarkMultiply100by200Times200by100(b *testing.B) {
	benchmarkMultiplyMatrix(100, 200, 200, 100, b)
}
