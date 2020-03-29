package pagerank

import "gonum.org/v1/gonum/mat"

func pageRank(P mat.Matrix, iterations int) *mat.Dense {
	r, _ := P.Dims()
	rank := Ones(r, 1)

	for i := 1; i <= iterations; i++ {
		rank.Mul(P.T(), rank)
	}
	return rank
}

func Ones(r, c int) *mat.Dense {
	data := make([]float64, r*c)
	for i := range data {
		data[i] = 1.0
	}
	m := mat.NewDense(r, c, data)
	return m
}
