package util

import "gonum.org/v1/gonum/mat"

func Add(x *mat.Dense, v ...float64) *mat.Dense {
	var dence mat.Dense
	for _, f := range v {
		dence.Apply(func(i, j int, v float64) float64 {
			return v + f
		}, x)
	}

	return &dence
}

func AddMatrix(xs ...*mat.Dense) *mat.Dense {
	if len(xs) == 0 {
		return nil
	}

	if len(xs) == 1 {
		return xs[0]
	}

	var result mat.Dense
	result.Add(xs[0], xs[1])
	for i := 2; i < len(xs); i++ {
		result.Add(&result, xs[i])
	}

	return &result

}

func Mul(x *mat.Dense, v ...float64) *mat.Dense {
	var dence mat.Dense
	for _, f := range v {
		dence.Apply(func(i, j int, v float64) float64 {
			return v * f
		}, x)
	}

	return &dence
}

func Div(x *mat.Dense, v ...float64) *mat.Dense {
	var dence mat.Dense
	for _, f := range v {
		dence.Apply(func(i, j int, v float64) float64 {
			return v / f
		}, x)
	}

	return &dence
}
