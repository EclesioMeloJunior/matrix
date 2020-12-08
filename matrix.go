package matrix

import (
	"math/rand"
	"reflect"
	"time"
)

// Matrix defines an [][]int64 type
type Matrix [][]int64

// Rows returns the quantity of rows in a matrix
func (m *Matrix) Rows() int {
	return len(*m)
}

// Cols returns the quantity of cols in a matrix
func (m *Matrix) Cols() int {
	if m.Rows() < 1 {
		return 0
	}

	return len((*m)[0])
}

// Transpose will transpose the matrix in O(r x c)
func (m *Matrix) Transpose() {
	newC := m.Rows()
	newR := m.Cols()

	newM := make([][]int64, newR)

	for i := 0; i < newR; i++ {
		newM[i] = make([]int64, newC)

		for j := 0; j < newC; j++ {
			newM[i][j] = (*m)[j][i]
		}
	}

	*m = newM
}

// Equals returns true if a matrix is equals another matrix
func (m *Matrix) Equals(o *Matrix) bool {
	if m.Rows() != o.Rows() || m.Cols() != o.Cols() {
		return false
	}

	return reflect.DeepEqual(m, o)
}

// Copy will a copy matrix based on root matrix
func (m *Matrix) Copy() *Matrix {
	newM := new(Matrix)
	*newM = make([][]int64, m.Rows())

	for i := 0; i < m.Rows(); i++ {
		(*newM)[i] = append((*newM)[i], (*m)[i]...)
	}

	return newM
}

// Zero creates a matrix R x C and all entries are 0
func Zero(r, c int) *Matrix {
	m := new(Matrix)
	*m = make([][]int64, r)

	for i := 0; i < r; i++ {
		(*m)[i] = make([]int64, c)

		for j := 0; j < c; j++ {
			(*m)[i][j] = 0
		}
	}

	return m
}

// RandInt creates a matrix R x C where all entries are choices from 0 - 100
func RandInt(r, c int) *Matrix {
	rand.Seed(time.Now().UnixNano())

	m := new(Matrix)
	*m = make([][]int64, r)

	for i := 0; i < r; i++ {
		(*m)[i] = make([]int64, c)
		for j := 0; j < c; j++ {
			(*m)[i][j] = rand.Int63n(100)
		}
	}

	return m
}
