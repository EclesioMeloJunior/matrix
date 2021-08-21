package matrix

func (m *Matrix) Determinant() (int64, error) {
	if !m.IsSquared() {
		return 0, ErrNotSquared
	}

	if m.Rows() == 1 {
		return m.Get(0, 0)
	}

	return calcDeterminant(m), nil
}

func calcDeterminant(m *Matrix) int64 {
	if m.Rows() == 2 {
		a1, _ := m.Get(0, 0)
		a2, _ := m.Get(1, 1)

		b1, _ := m.Get(0, 1)
		b2, _ := m.Get(1, 0)
		return (a1 * a2) - (b1 * b2)
	}

	var det int64

	for i := 0; i < m.Cols(); i++ {
		subset := GetDeterminantSubset(m, i)
		v, _ := m.Get(0, i)
		currDet := v * calcDeterminant(subset)

		if i&1 == 0 {
			det = det + currDet
		} else {
			det = det - currDet
		}
	}

	return det
}

func GetDeterminantSubset(m *Matrix, col int) *Matrix {
	if m.Rows() <= 2 {
		return nil
	}

	newM := new(Matrix)
	*newM = make([][]int64, m.Rows()-1)

	subsetCols := m.Cols() - 1
	for i := 1; i < m.Rows(); i++ {
		(*newM)[i-1] = make([]int64, subsetCols)
		copy((*newM)[i-1][:], (*m)[i][:col])
		copy((*newM)[i-1][col:], (*m)[i][col+1:])
	}

	return newM
}
