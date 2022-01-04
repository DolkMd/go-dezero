package util

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Dence1D(dence ...float64) *mat.Dense {
	return mat.NewDense(1, len(dence), dence)
}

func Dence2D(dence [][]float64) *mat.Dense {
	tmpDence := []float64{}
	for _, denceArr := range dence {
		tmpDence = append(tmpDence, denceArr...)
	}
	return mat.NewDense(len(dence[0]), len(dence), tmpDence)
}

func EachDence(dence *mat.Dense, fn func(float64)) {
	for i := 0; i < dence.RawMatrix().Rows; i++ {
		for j := 0; j < dence.RawMatrix().Cols; j++ {
			fn(dence.At(i, j))
		}
	}
}

func Allclose(a, b *mat.Dense, rtol, atol *float64) bool {
	if rtol == nil {
		rtol = new(float64)
		*rtol = 1e-5
	}
	if atol == nil {
		atol = new(float64)
		*atol = 1e-8
	}

	for i := 0; i < a.RawMatrix().Rows; i++ {
		for j := 0; j < a.RawMatrix().Cols; j++ {
			aij := a.At(i, j)
			bij := b.At(i, j)
			if math.Abs(aij-bij) > (*atol + (*rtol)*math.Abs(bij)) {
				return false
			}
		}
	}
	return true
}

func FullOne(n int) []float64 {
	r := make([]float64, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}
	return r
}
