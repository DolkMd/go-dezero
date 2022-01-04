package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"

	"gonum.org/v1/gonum/mat"
)

type (
	Div interface {
		Layer
	}
	div struct {
		core.Function
	}
)

func NewDiv() Div {
	instance := new(mul)
	instance.Function = core.NewFunction(instance.Forward, instance.Backward)
	return instance
}

func (d *div) Forward(variables ...core.Variable) core.Variables {
	var y mat.Dense
	y.DivElem(variables[0].Data(), variables[1].Data())
	return core.NewVariables(&y)
}

func (d *div) Backward(variables ...core.Variable) core.Variables {
	panic("")
	// x0, x1 := d.Inputs()[0].Data(), d.Inputs()[1].Data()
	// y1 := function.Div(variables[0], variables[1])
	// y2 := function.Mul(variables[0],
	// 	function.Neg(
	// 		function.Div(
	// 			x0, function.Pow(x1, 2.0)[0],
	// 		)[0],
	// 	)[0],
	// )
	// return []core.Variable{y1[0], y2[0]}
}
