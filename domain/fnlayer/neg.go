package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type (
	Neg interface {
		Layer
	}
	neg struct {
		core.Function
	}
)

func NewNeg() Neg {
	instance := new(neg)
	instance.Function = core.NewFunction(instance.Forward, instance.Backward)
	return instance
}

func (m *neg) Forward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Apply(func(i, j int, v float64) float64 {
		return v * -1
	}, variables[0].Data())
	return core.NewVariables(&y)
}

func (m *neg) Backward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Apply(func(i, j int, v float64) float64 {
		return v * -1
	}, variables[0].Data())
	return core.NewVariables(&y)
}
