package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type (
	Add interface {
		Layer
	}
	add struct {
		core.Function
	}
)

func NewAdd() Add {
	instance := new(add)
	instance.Function = core.NewFunction(instance.Forward, instance.Backward)
	return instance
}

func (a *add) Forward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.Add(variables[0].Data(), variables[1].Data())
	return core.NewVariables(&y)
}

func (s *add) Backward(variables ...core.Variable) core.Variables {
	return []core.Variable{variables[0], variables[0]}
}

func (s *add) Apply(xs ...core.Variable) core.Variables {
	if len(xs) != 2 {
		panic("add apply only two args")
	}

	return s.Function.Apply(xs...)
}
