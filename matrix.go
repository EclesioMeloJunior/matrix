package matrix

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

// ErrMatrixOfDifferentSizes error when a matrix has not the same rows or cols of another
var ErrMatrixOfDifferentSizes = errors.New("Matrix have not same rows or same cols")
var ErrNotSquared = errors.New("matrix must be squared to perform the operation")
var ErrOutOfTheMatrix = errors.New("matrix does not contains index")

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

func (m *Matrix) Get(i, j int) (int64, error) {
	if i >= m.Rows() || j >= m.Cols() {
		return 0, ErrOutOfTheMatrix
	}

	return (*m)[i][j], nil
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

// Copy O(r) will a copy matrix based on root matrix
func (m *Matrix) Copy() *Matrix {
	newM := new(Matrix)
	*newM = make([][]int64, m.Rows())

	for i := 0; i < m.Rows(); i++ {
		(*newM)[i] = append((*newM)[i], (*m)[i]...)
	}

	return newM
}

// IsSquared returns if the matrix has the same number of rows and columns
func (m *Matrix) IsSquared() bool {
	return m.Rows() == m.Cols()
}

// Sum O(r x c) will return a new matrix with the sum of the 2 matrix
func (m *Matrix) Sum(o *Matrix) (*Matrix, error) {
	if m.Rows() != o.Rows() || m.Cols() != o.Cols() {
		return nil, ErrMatrixOfDifferentSizes
	}

	newM := new(Matrix)
	*newM = make([][]int64, m.Rows())

	for i := 0; i < m.Rows(); i++ {
		(*newM)[i] = make([]int64, m.Cols())

		for j := 0; j < m.Cols(); j++ {
			(*newM)[i][j] = (*m)[i][j] + (*o)[i][j]
		}
	}

	return newM, nil
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
			(*m)[i][j] = rand.Int63n(255)
		}
	}

	return m
}
