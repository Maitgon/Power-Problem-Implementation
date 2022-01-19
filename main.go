package main

import (
	"fmt"
	"implementation/power1"
)

func main() {
	input := power1.CreateInstancePower1(14, []int{1, 2, 3, 4, 5, 6})
	pow1, sol1 := power1.SolvePower1(input)

	fmt.Println(pow1, sol1)

	input2 := power1.CreateInstancePower1(460, []int{30, 45, 12, 16, 19, 34, 56, 19, 14, 10, 87, 45, 26, 38, 57})
	pow2, sol2 := power1.SolvePower1(input2)

	fmt.Println(pow2, sol2)
}
