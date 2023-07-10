package main

import "fmt"

// SumInts add together the value of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the value of m.
func SumFloats (m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main () {
	// Initialize a map of strings to int64.
	ints := map[string]int64{
		"one": 1,
		"two": 2,
	}

	// Initialize a map of strings to float64.
	floats := map[string]float64{
		"pi": 3.14,
		"e": 2.718,
	}

	fmt.Printf("Non-Generic Sums ; Ints: %d, Floats: %f\n", SumInts(ints), SumFloats(floats))
}