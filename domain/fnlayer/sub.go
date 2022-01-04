package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type (
	Sub interface{ Layer }
	sub struct{ core.Function }
)

func NewSub() Sub {
	instance := new(sub)
	instance.Function = core.NewFunction(instance.Forward, instance.Backward)
	return instance
}

func (m *sub) Forward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Sub(variables[0].Data(), variables[1].Data())
	return core.NewVariables(&y)
}

func (m *sub) Backward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Apply(func(i, j int, v float64) float64 {
		return v * -1
	}, variables[0].Data())
	return core.NewVariables(variables[0].Data(), &y)
}
