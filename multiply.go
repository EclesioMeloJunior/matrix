package matrix

// MultplyByScalar will multply a matrix by a single value
func (m *Matrix) MultplyByScalar(v int64) {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			(*m)[i][j] = (*m)[i][j] * v
		}
	}
}
