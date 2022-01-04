package function

import (
	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/DolkMd/go-dezero/domain/fnlayer"
)

func Square(xs ...core.Variable) core.Variables {
	return fnlayer.NewSquare().Apply(xs...)
}
func Exp(xs ...core.Variable) core.Variables {
	return fnlayer.NewExp().Apply(xs...)
}
func Add(xs ...core.Variable) core.Variables {
	return fnlayer.NewAdd().Apply(xs...)
}
func Mul(xs ...core.Variable) core.Variables {
	return fnlayer.NewMul().Apply(xs...)
}
func Neg(xs ...core.Variable) core.Variables {
	return fnlayer.NewNeg().Apply(xs...)
}
func Sub(xs ...core.Variable) core.Variables {
	return fnlayer.NewSub().Apply(xs...)
}
func Div(xs ...core.Variable) core.Variables {
	return fnlayer.NewDiv().Apply(xs...)
}
func Composite(fns ...func(...core.Variable) core.Variables) func(...core.Variable) core.Variables {
	return func(v ...core.Variable) core.Variables {
		for _, fn := range fns {
			v = fn(v...)
		}

		return v
	}
}
