// power1/power1.go

package power1

import (
	"implementation/utils"
	"math"
)

type Power1 struct {
	k int
	s []int
}

func CreateInstancePower1(k int, s []int) Power1 {
	return Power1{k, s}
}

func Combinations1(p Power1) int {

	// The bits of picked are the elements we choose
	// 5 = 00101: we choose 1st and 3rd.
	var picked int64

	if len(p.s) > 64 {
		return 0
	}

	var sol int
	for ; float64(picked) < math.Pow(2, float64(len(p.s))); picked++ {
		aux := 0
		for i := 0; i < len(p.s); i++ {
			if utils.IsKthBitSet(picked, i) {
				aux += p.s[i]
			}
		}
		if aux >= p.k {
			sol++
		}
	}

	return sol

}

func SolvePower1(p Power1) ([]int, int) {

	// The bits of picked are the elements we choose
	// 5 = 00101: we choose 1st and 3rd.
	var picked int64

	if len(p.s) > 64 {
		return []int{}, 0
	}

	powers := make([]int, len(p.s))

	var sol int
	for ; float64(picked) < math.Pow(2, float64(len(p.s))); picked++ {
		aux := 0

		for i := 0; i < len(p.s); i++ {
			if utils.IsKthBitSet(picked, i) {
				aux += p.s[i]
			}
		}

		if aux >= p.k {
			sol++
			for i := 0; i < len(p.s); i++ {
				if utils.IsKthBitSet(picked, i) {
					powers[i]++
				}
			}
		}
	}

	return powers, sol

}
