package fnlayer

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type Exp interface {
	Layer
}

type exp struct {
	core.Function
}

func NewExp() Exp {
	instance := new(exp)
	instance.Function = core.NewFunction(
		instance.Forward,
		instance.Backward,
	)
	return instance
}

func (*exp) Forward(variables ...core.Variable) core.Variables {
	var result mat.Dense
	result.Apply(func(i, j int, v float64) float64 {
		return math.Exp(v)
	}, variables[0].Data())

	return core.NewVariables(&result)
}

func (e *exp) Backward(variables ...core.Variable) core.Variables {
	x := *e.Inputs()[0].Data()
	x.Apply(func(i, j int, v float64) float64 {
		return math.Exp(v)
	}, &x)
	x.Mul(&x, variables[0].Data())

	return core.NewVariables(&x)
}
