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
