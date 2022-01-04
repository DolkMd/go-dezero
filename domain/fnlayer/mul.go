package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type (
	Mul interface {
		Layer
	}
	mul struct {
		core.Function
	}
)

func NewMul() Mul {
	instance := new(mul)
	instance.Function = core.NewFunction(instance.Forward, instance.Backward)
	return instance
}

func (m *mul) Forward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Mul(variables[0].Data(), variables[1].Data())

	return core.NewVariables(&y)
}

func (m *mul) Backward(variables ...core.Variable) core.Variables {
	x0, x1 := m.Inputs()[0].Data(), m.Inputs()[1].Data()
	var y1, y2 mat.Dense
	y1.Mul(variables[0].Data(), x1)
	y2.Mul(variables[0].Data(), x0)
	return core.NewVariables(&y1, &y2)
}
