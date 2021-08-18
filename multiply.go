package matrix

import (
	"errors"
)

// MultplyByScalar will multply a matrix by a single value
func (m *Matrix) MultplyByScalar(v int64) {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			(*m)[i][j] = (*m)[i][j] * v
		}
	}
}

// Multiply will multiply 2 matrices and return a new matrix
func (m *Matrix) Multiply(o *Matrix) (*Matrix, error) {
	if m.Cols() != o.Rows() {
		return nil, errors.New("the matrix has not the same rows and columns")
	}

	r := new(Matrix)
	*r = make([][]int64, m.Rows())

	for i := 0; i < m.Rows(); i++ {
		(*r)[i] = make([]int64, o.Cols())

		for j := 0; j < o.Cols(); j++ {

			for k := 0; k < o.Rows(); k++ {
				v := (*m)[i][k] * (*o)[k][j]
				(*r)[i][j] = (*r)[i][j] + v
			}
		}
	}

	return r, nil
}
