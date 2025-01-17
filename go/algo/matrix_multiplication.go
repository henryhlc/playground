package algo

import (
	"fmt"
	"strings"
)

type Matrix struct {
	n, m int
	// Column-major
	data []float64
}

func (m Matrix) index(i, j int) int {
	return j*m.n + i
}

func (m Matrix) Print() {
	for i := range m.n {
		row := make([]string, m.m)
		for j := range m.m {
			row = append(row, fmt.Sprintf("%v", m.At(i, j)))
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func (m Matrix) Set(i, j int, v float64) {
	m.data[m.index(i, j)] = v
}

func (m Matrix) At(i, j int) float64 {
	return m.data[m.index(i, j)]
}

// Add the product of A, B to the matrix m.
// Noop if the dimensions do not match.
func (m Matrix) MultiplyAddIJP(A, B Matrix) {
	if m.n != A.n || m.m != B.m || A.m != B.n {
		return
	}
	for i := range A.n {
		for j := range B.m {
			for p := range B.n {
				m.Set(i, j, m.At(i, j)+A.At(i, p)*B.At(p, j))
			}
		}
	}
}
