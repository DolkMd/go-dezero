package fnlayer

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type Square interface {
	Layer
}

type square struct {
	core.Function
}

func NewSquare() Square {
	instance := new(square)
	instance.Function = core.NewFunction(
		instance.Forward,
		instance.Backward,
	)
	return instance
}

func (*square) Forward(variables ...core.Variable) core.Variables {
	var result mat.Dense
	result.Apply(func(i, j int, v float64) float64 {
		return math.Pow(v, 2)
	}, variables[0].Data())

	return core.NewVariables(&result)
}

func (s *square) Backward(variables ...core.Variable) core.Variables {
	var result mat.Dense
	result.Apply(func(i, j int, v float64) float64 {
		return v * 2
	}, s.Inputs()[0].Data())
	result.Mul(&result, variables[0].Data())

	return core.NewVariables(&result)
}
