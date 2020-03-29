package pagerank

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func pageRank(P mat.Matrix, tolerance float64) mat.Vector {
	r, _ := P.Dims()
	data := make([]float64, r)
	for i := range data {
		data[i] = float64(1 / r)
	}
	rank := mat.NewVecDense(r, data)
	rankPre := mat.NewVecDense(r, data)

	for {
		rank, rankPre = rankPre, rank
		rank.MulVec(P, rankPre)
		if vecDiff(rank, rankPre) < tolerance {
			break
		}
	}

	return rank
}

func vecDiff(a, b mat.Vector) float64 {
	if a.Len() != b.Len() {
		panic("Two vectors must have same length!")
	}
	var sum float64
	for i := 0; i < a.Len(); i++ {
		sum += math.Abs(b.AtVec(i) - a.AtVec(i))
	}
	return sum

}

func Ones(r, c int) *mat.Dense {
	data := make([]float64, r*c)
	for i := range data {
		data[i] = 1.0
	}
	m := mat.NewDense(r, c, data)
	return m
}
