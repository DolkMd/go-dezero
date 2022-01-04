package core

import (
	"gonum.org/v1/gonum/mat"
)

type Variables []Variable

func NewVariables(data ...*mat.Dense) Variables {
	result := make([]Variable, 0, len(data))
	for _, v := range data {
		result = append(result, NewVariable(v))
	}

	return result
}

func (v Variables) First() Variable {
	if len(v) == 0 {
		panic("not found first variable")
	}

	return v[0]
}

func (v Variables) Grads() []*mat.Dense {
	result := []*mat.Dense{}
	for _, variable := range v {
		if variable.Grad() != nil {
			result = append(result, variable.Grad())
		}
	}

	return result
}

func (v Variables) Generations() []int {
	result := []int{}
	for _, variable := range v {
		result = append(result, variable.Generation())
	}

	return result
}

func (v Variables) ToDataList() []*mat.Dense {
	result := make([]*mat.Dense, 0, len(v))
	for _, v := range v {
		result = append(result, v.Data())
	}

	return result
}
