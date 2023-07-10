package main

import "fmt"
// Declare the Number interface type to use as a type constraint.
// Declare a union of int64 and float64 inside the interface.
type Number interface {
	int64 | float64
}

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

// sumIntsOrFloats adds the values of map m. It supports both int64 and float64 as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. It supports both int64 and float64 as types for map values.
func SumNumbers[K comparable, V Number](m map[K]V) V { 	
	var s V
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
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
}