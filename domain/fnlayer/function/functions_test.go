package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/DolkMd/go-dezero/domain/fnlayer/function"
	"github.com/DolkMd/go-dezero/domain/util"
)

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}
func TestSquare(t *testing.T) {
	t.Run("success: test Square model", func(t *testing.T) {
		x := core.NewVariable(util.Dence1D(2.0))
		y := function.Square(x)
		assert.Equal(t, util.Dence1D(4.0), y.First().Data())
	})

	t.Run("success: test Square model backward", func(t *testing.T) {
		x := core.NewVariable(util.Dence1D(3.0))
		y := function.Square(x)
		y.First().Backward()

		assert.Equal(t, x.Grad(), util.Dence1D(6.0))
	})

	t.Run("success: Square model gradient", func(t *testing.T) {
		x := core.NewVariable(util.Dence1D(random(0, 1)))
		y := function.Square(x)
		y.First().Backward()

		gard := core.NumericalDiff(function.Square, x, nil)
		flag := util.Allclose(x.Grad(), gard, nil, nil)
		t.Log(fmt.Sprintf("grad: %v, x-grad: %v", gard.At(0, 0), x.Grad().At(0, 0)))
		assert.True(t, flag)
	})

}
