package core

import (
	"sort"

	"github.com/DolkMd/go-dezero/domain/util"

	"gonum.org/v1/gonum/mat"
)

type (
	Variable interface {
		Data() *mat.Dense

		Grad() *mat.Dense
		SetGrad(*mat.Dense)
		ClearGrad()

		Creator() Function
		SetCreator(Function)

		NDim() (int, int)
		Size() int

		Generation() int

		Backward(...VariableBackfowardOpt)
	}
	variable struct {
		data       *mat.Dense
		gard       *mat.Dense
		creator    Function
		generation int
		name       string
	}
)

func NewVariable(data *mat.Dense, opts ...VariableOpt) Variable {
	instance := &variable{data: data}
	for _, opt := range opts {
		opt(instance)
	}
	return instance
}

func (v *variable) Data() *mat.Dense {
	return v.data
}

func (v *variable) Grad() *mat.Dense {
	return v.gard
}

func (v *variable) SetGrad(dence *mat.Dense) {
	v.gard = dence
}

func (v *variable) ClearGrad() {
	v.gard = nil
}

func (v *variable) Creator() Function {
	return v.creator
}

func (v *variable) SetCreator(creator Function) {
	v.generation = creator.Generation() + 1
	v.creator = creator
}

func (v *variable) Generation() int {
	return v.generation
}

func (v *variable) Backward(opts ...VariableBackfowardOpt) {
	opt := varBackfowardOpts{}
	for _, o := range opts {
		o(&opt)
	}

	if v.gard == nil {
		copiedData := mat.DenseCopyOf(v.data)
		copiedData.Apply(func(i, j int, v float64) float64 { return 1.0 }, copiedData)
		v.gard = copiedData
	}

	creators := []Function{}
	seenSet := util.Set()
	addFunc := func(f Function) {
		if !seenSet.Find(f) {
			creators = append(creators, f)
			seenSet.Add(f)
			sort.Slice(creators, func(i, j int) bool {
				return creators[i].Generation() > creators[j].Generation()
			})
		}
	}
	addFunc(v.creator)

	for len(creators) > 0 {
		creator := creators[0]
		creators = creators[1:]
		gys := creator.Outputs().Grads()
		gxs := creator.Backward(NewVariables(gys...)...)
		for i := range creator.Inputs() {
			x := creator.Inputs()[i]
			gx := gxs[i]
			if x.Grad() == nil {
				x.SetGrad(gx.Data())
			} else {
				x.SetGrad(util.AddMatrix(x.Grad(), gx.Data()))
			}
			if x.Creator() != nil {
				addFunc(x.Creator())
			}
		}
		if !opt.retainGradGrad {
			for _, y := range creator.Outputs() {
				y.SetGrad(nil)
			}
		}
	}
}

func (v *variable) NDim() (int, int) {
	return v.data.Dims()
}
func (v *variable) Size() int {
	r, c := v.NDim()
	return r * c
}

// variable options

type (
	VariableOpt func(*variable)
	varOpts     struct{}

	VariableBackfowardOpt func(*varBackfowardOpts)
	varBackfowardOpts     struct {
		retainGradGrad bool
	}
)

func VarOpts() varOpts {
	return varOpts{}
}

func (varOpts) Name(name string) VariableOpt {
	return func(v *variable) {
		v.name = name
	}
}

func VarBackfowardOpt() varBackfowardOpts {
	return varBackfowardOpts{}
}

func (varBackfowardOpts) RetainGradGrad(retainGradGrad bool) VariableBackfowardOpt {
	return func(opt *varBackfowardOpts) {
		opt.retainGradGrad = retainGradGrad
	}
}
