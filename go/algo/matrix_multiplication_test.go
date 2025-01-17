package algo

import (
	"math"
	"math/rand/v2"
	"testing"
)

func TestMultiplyAdd(t *testing.T) {
	C := Matrix{
		n:    3,
		m:    3,
		data: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	A := Matrix{
		n:    3,
		m:    2,
		data: []float64{1, 2, 3, 4, 5, 6},
		// 1 4
		// 2 5
		// 3 6
	}
	B := Matrix{
		n:    2,
		m:    3,
		data: []float64{1, 2, 3, 4, 5, 6},
		// 1 3 5
		// 2 4 6
	}
	// 9 19 29
	// 12 26 40
	// 15 33 51
	expectedData := []float64{9, 12, 15, 19, 26, 33, 29, 40, 51}
	C.MultiplyAddIJP(A, B)
	for i := range len(C.data) {
		if math.Abs(C.data[i]-expectedData[i]) > 1e-5 {
			t.Errorf("Actual %v != expected %v", C.data, expectedData)
			break
		}
	}
}

func randomSquareMatrix(n int) Matrix {
	m := Matrix{
		n:    n,
		m:    n,
		data: make([]float64, n*n),
	}
	for i := range n * n {
		m.data[i] = rand.Float64()
	}
	return m
}

func BenchmarkMultiplyAdd(b *testing.B) {
	CIJP := randomSquareMatrix(100)
	CJPI := randomSquareMatrix(100)
	CPIJ := randomSquareMatrix(100)
	A := randomSquareMatrix(100)
	B := randomSquareMatrix(100)
	b.Run("IJP", func(b *testing.B) {
		for range b.N {
			CIJP.MultiplyAddIJP(A, B)
		}
	})
	b.Run("JPI", func(b *testing.B) {
		for range b.N {
			CJPI.MultiplyAddJPI(A, B)
		}
	})
	b.Run("PIJ", func(b *testing.B) {
		for range b.N {
			CPIJ.MultiplyAddPIJ(A, B)
		}
	})
}
