package pagerank

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestOnes(t *testing.T) {
	m := Ones(10, 3)
	mat.Formatted(m)
}
