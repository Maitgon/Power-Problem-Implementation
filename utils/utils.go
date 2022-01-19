// utils/utils.go

package utils

func IsKthBitSet(n int64, k int) bool {
	if n&(1<<k) != 0 {
		return true
	} else {
		return false
	}
}

func Sum(s []int) int {
	sol := 0
	for _, val := range s {
		sol += val
	}
	return sol
}
