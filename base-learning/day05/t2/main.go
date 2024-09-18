package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {

	fmt.Println("hello world...")

	m1 := Max[int]([]int{40, -8, 15})
	m2 := Max[float64]([]float64{4.1, -8.1, 15.1})

	fmt.Println(m1, m2)
}

func Max[T constraints.Ordered](s []T) T {

	var zero T

	if len(s) == 0 {
		return zero
	}

	var max T
	max = s[0]

	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
