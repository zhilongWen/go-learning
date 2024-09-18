package main

import "fmt"

func main() {

	println("hello world...")

	m := MaxInt([]int{1, 2, 3, 4, 5})
	fmt.Println(m)

	m = MaxInt(nil)
	fmt.Println(m)

	m1 := MaxFloat([]float64{1.1, 1.2, 1.3, 1.4, 5.0})
	fmt.Println(m1)
}

func MaxInt(s []int) int {
	if len(s) == 0 {
		return 0
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

func MaxFloat(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
