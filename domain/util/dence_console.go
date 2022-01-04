package util

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func PrintDence(x *mat.Dense) {
	if x != nil {
		fa := mat.Formatted(x, mat.Prefix(" "), mat.Squeeze())
		fmt.Printf("%v\n", fa)
	}

}

func PrintDences(dences ...*mat.Dense) {
	for _, dence := range dences {
		PrintDence(dence)
	}
}
