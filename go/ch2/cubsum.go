package ch2

import (
	"math"
)

func CubSumAB(n int) [][3]int {
	minpq := NewIndexMinPQ(n+1)
	ij := make([]int, n+1)

	// iterating i from 0 to n to create all tuples (i^3+j^3, i, j=0)
	for i := range ij {
		// j=0 for the first n+1 pairs
		ij[i] = 0

		key := int(math.Pow(float64(i), 3.0))
		minpq.Insert(i, key)
	}

	// building all combinations for (i,j), i, j <= n
	res := make([][3]int, 0)
	for !minpq.IsEmpty() {
		min := minpq.Min()
		i := minpq.DelMin()
		j := ij[i]
		res = append(res, [...]int{min, i, j})

		if j < n {
			ij[i] = j+1
			key := int(math.Pow(float64(i), 3.0) + math.Pow(float64(j+1), 3.0))
			minpq.Insert(i, key)
		}
	}
	return res
}