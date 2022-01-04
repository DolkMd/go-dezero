package core

import (
	conf "github.com/DolkMd/go-dezero/domain/core/config"
	"github.com/DolkMd/go-dezero/domain/util"

	xmath "github.com/fcfcqloow/go-advance/math"
	"gonum.org/v1/gonum/mat"
)

type (
	Forward  func(...Variable) Variables
	Backward func(...Variable) Variables
)

type (
	Function interface {
		Apply(...Variable) Variables
		Inputs() Variables
		Outputs() Variables
		Backward(...Variable) Variables
		Forward(...Variable) Variables
		Generation() int
	}
	function struct {
		forward    Forward
		backward   Backward
		inputs     Variables
		outputs    Variables
		generation int
	}
)

func NewFunction(forward Forward, backward Backward) Function {
	return &function{
		forward:  forward,
		backward: backward,
	}
}

func (f *function) Apply(inputs ...Variable) Variables {
	xs := Variables(inputs)
	outputs := f.forward(xs...)

	if conf.Config.EnableBackprop {
		f.generation = xmath.MaxInt(Variables(inputs).Generations()...)
		for _, output := range outputs {
			output.SetCreator(f)
		}
	}

	f.inputs = inputs
	f.outputs = outputs

	return outputs
}

func (f *function) Inputs() Variables {
	return f.inputs
}

func (f *function) Outputs() Variables {
	return f.outputs
}

func (f *function) Backward(dence ...Variable) Variables {
	return f.backward(dence...)
}

func (f *function) Forward(dence ...Variable) Variables {
	return f.forward(dence...)
}

func (f *function) Generation() int {
	return f.generation
}

func NumericalDiff(f func(...Variable) Variables, x Variable, eps *float64) *mat.Dense {
	if eps == nil {
		eps = new(float64)
		*eps = 1e-4
	}

	x0 := NewVariable(util.Add(x.Data(), -*eps))
	x1 := NewVariable(util.Add(x.Data(), *eps))
	y0 := f(x0)
	y1 := f(x1)

	var result mat.Dense
	result.Sub(y1[0].Data(), y0[0].Data())
	result.Apply(func(i, j int, v float64) float64 {
		return v / (2 * (*eps))
	}, &result)

	return &result
}
