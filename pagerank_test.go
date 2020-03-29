package pagerank

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func TestOnes(t *testing.T) {
	m := Ones(10, 3)
	matPrint(m)
}
