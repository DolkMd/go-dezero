package fnlayer

import (
	"github.com/DolkMd/go-dezero/domain/core"
)

type Layer interface {
	Forward(...core.Variable) core.Variables
	Backward(...core.Variable) core.Variables
	Apply(...core.Variable) core.Variables
	Inputs() core.Variables
}

func Composite(models ...Layer) func(...core.Variable) core.Variables {
	return func(v ...core.Variable) core.Variables {
		for _, model := range models {
			v = model.Apply(v...)
		}
		return v
	}
}
