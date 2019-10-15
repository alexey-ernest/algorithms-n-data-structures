package ch2

import (
	//"fmt"
	"math"
	"testing"
)

func TestCubSumABBasic(t *testing.T) {
	res := CubSumAB(2) // 0,1,2
	exp := [][3]int{
		[...]int{0, 0, 0}, 
		[...]int{1, 1, 0}, 
		[...]int{1, 0, 1},
		[...]int{2, 1, 1},
		[...]int{8, 2, 0},
		[...]int{8, 0, 2},
		[...]int{9, 2, 1},
		[...]int{9, 1, 2},
		[...]int{16, 2, 2},
	}

	for i := range res {
		found := false
		for j := range exp {
			if res[i] == exp[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%+v is not found in expected result set", res[i])
			return
		}
	}
}


func TestCubSumABLarge(t *testing.T) {
	n := 100
	res := CubSumAB(n)
	exp := [][3]int{}
	for i := 0; i <= n; i+= 1 {
		for j := 0; j <= n; j+=1 {
			exp = append(exp, [...]int{int(math.Pow(float64(i), 3) + math.Pow(float64(j), 3)), i, j})
		}
	}

	lastsum := 0
	for i := range res {
		found := false
		if res[i][0] < lastsum {
			t.Errorf("non-decreasing condition failed")
			return
		}
		lastsum = res[i][0]

		for j := range exp {
			if res[i] == exp[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%+v is not found in expected result set", res[i])
			return
		}
	}
}

func TestCubSumABCDBasic(t *testing.T) {
	n := 20
	res := CubSumAB(n)
	exp := [][3]int{}
	for i := 0; i <= n; i+= 1 {
		for j := 0; j <= n; j+=1 {
			exp = append(exp, [...]int{int(math.Pow(float64(i), 3) + math.Pow(float64(j), 3)), i, j})
		}
	}

	lastsum := 0
	lastsumpairs := [][2]int{}
	abcdres := [][5]int{}
	for i := range res {
		found := false
		if res[i][0] < lastsum {
			t.Errorf("non-decreasing condition failed")
			return
		}

		if res[i][0] != lastsum {
			lastsum = res[i][0]

			if len(lastsumpairs) > 0 {
				for i := range lastsumpairs {
					for j := range lastsumpairs {
						abcd := [5]int{lastsum, lastsumpairs[i][0], lastsumpairs[i][1], lastsumpairs[j][0], lastsumpairs[j][1]}
						//fmt.Printf("%+v\n", abcd)
						abcdres = append(abcdres, abcd)
					}
				}
			}
			lastsumpairs = [][2]int{}
		}
		lastsumpairs = append(lastsumpairs, [2]int{res[i][1], res[i][2]})

		for j := range exp {
			if res[i] == exp[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%+v is not found in expected result set", res[i])
			return
		}
	}
}