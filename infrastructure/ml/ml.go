package ml

import (
	"math/rand"
	"time"

	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/DolkMd/go-dezero/domain/fnlayer"
	. "github.com/DolkMd/go-dezero/domain/fnlayer/function"
	"github.com/DolkMd/go-dezero/domain/util"
)

// func Script1() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a)
// 	y := C.Apply(b)
// 	util.PrintDence(y.Data())
// }

func Script2() {
	f := fnlayer.NewSquare()
	x := core.NewVariable(util.Dence1D(2.0))
	dy := core.NumericalDiff(f.Apply, x, nil)
	util.PrintDence(dy)

}

// func Script3() {
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	dy := core.NumericalDiff(fnlayer.Composite(
// 		fnlayer.NewSquare(),
// 		fnlayer.NewExp(),
// 		fnlayer.NewSquare(),
// 	), x, nil)
// 	util.PrintDence(dy)
// }

// func Script4() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a)
// 	y := C.Apply(b)
// 	util.PrintDence(y.Data()) // 1.648721270700128
// 	y.SetGrad(util.Dence1D(1.0))
// 	b.SetGrad(C.Backward(y.Grad()))
// 	a.SetGrad(B.Backward(b.Grad()))
// 	x.SetGrad(A.Backward(a.Grad()))
// 	util.PrintDence(x.Grad()) // 3.297442541400256
// }

// func Script5() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a)
// 	y := C.Apply(b)
// 	util.PrintDence(y.Data()) // 1.648721270700128
// 	y.SetGrad(util.Dence1D(1.0))
// 	Cfn := y.Creator()
// 	b = Cfn.Input()
// 	b.SetGrad(Cfn.Backward(y.Grad()))
// 	Bfn := b.Creator()
// 	a = Bfn.Input()
// 	a.SetGrad(Bfn.Backward(b.Grad()))
// 	Afn := a.Creator()
// 	x = Afn.Input()
// 	x.SetGrad(Afn.Backward(a.Grad()))
// 	util.PrintDence(x.Grad()) // 3.297442541400256
// }

// func Script6() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a)
// 	y := C.Apply(b)
// 	util.PrintDence(y.Data()) // 1.648721270700128
// 	y.SetGrad(util.Dence1D(1.0))
// 	y.Backward()
// 	util.PrintDence(x.Grad()) // 3.297442541400256
// }

// func Script7() {
// 	x := core.NewVariable(util.Dence1D(0.5))
// 	a := Square(x)
// 	b := Exp(a)
// 	y := Square(b)
// 	y.Backward()
// 	util.PrintDence(x.Grad()) // 3.297442541400256
// }

func Script8() {
	x1, x2 := util.Dence1D(2), util.Dence1D(3)
	y := Add(core.NewVariables(x1, x2)...)
	util.PrintDence(y.First().Data())
}

func Script9() {
	x := core.NewVariable(util.Dence1D(2.0))
	y := core.NewVariable(util.Dence1D(3.0))
	z := Add(Square(x).First(), Square(y).First())
	z.First().Backward()

	util.PrintDences(z.First().Data())
	util.PrintDences(x.Grad())
	util.PrintDences(y.Grad())
}

func Script10() {
	x := core.NewVariable(util.Dence1D(3.0))
	y := Add(x, x)
	y.First().Backward()
	util.PrintDences(x.Grad()) // 2

	x.ClearGrad()
	y = Add(Add(x, x).First(), x)
	y.First().Backward()
	util.PrintDence(x.Grad()) // 3
}

func Script11() {
	x := core.NewVariable(util.Dence1D(2.0))       // x = 2
	a := Square(x).First()                         // x^2 = 4
	y := Add(Square(a).First(), Square(a).First()) // x^2^2 + x^2^2 = 2(x^4)= 32
	y.First().Backward()                           // y' = 8x^3 = 64

	util.PrintDence(y.First().Data())
	util.PrintDence(x.Grad())
}

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}
func Script12() {
	list := []float64{}
	for i := 0; i < 10000; i++ {
		list = append(list, random(0, 10000))
	}

	for i := 0; i < 10; i++ {
		x := core.NewVariable(util.Dence1D(list...))
		y := Composite(Square, Square, Square)(x)
		util.PrintDences(y.ToDataList()...)

	}
}

func Script13() {
	a := core.NewVariable(util.Dence1D(3.0))
	b := core.NewVariable(util.Dence1D(2.0))
	c := core.NewVariable(util.Dence1D(1.0))
	y := Add(Mul(a, b).First(), c)
	y.First().Backward()
	util.PrintDences(y.First().Data(), a.Grad(), b.Grad())

	x := core.NewVariable(util.Dence1D(2.0))
	y = Neg(x)
	util.PrintDence(y.First().Data())

	x = core.NewVariable(util.Dence1D(2.0))
	y1 := Sub(core.NewVariable(util.Dence1D(2.0)), x)
	y2 := Sub(x, core.NewVariable(util.Dence1D(1.0)))
	util.PrintDences(y1.First().Data(), y2.First().Data())
}
