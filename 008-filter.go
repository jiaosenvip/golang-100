package main

import "fmt"

func FilterInt(arr []int, f func(int) bool) []int {
	arf := make([]int, 0)
	for _, v := range arr {
		if f(v) {
			arf = append(arf, v)
		}
	}
	return arf
}
func FilterFloat64(arr []float64, f func(float64) bool) []float64 {
	arf := make([]float64, 0)
	for _, v := range arr {
		if f(v) {
			arf = append(arf, v)
		}
	}
	return arf
}
func FilterBool(arr []bool, f func(bool) bool) []bool {
	arf := make([]bool, 0)
	for _, v := range arr {
		if f(v) {
			arf = append(arf, v)
		}
	}
	return arf
}
func FilterString(arr []string, f func(string) bool) []string {
	arf := make([]string, 0)
	for _, v := range arr {
		if f(v) {
			arf = append(arf, v)
		}
	}
	return arf
}
func main() {
	intCheck := func(x int) bool { return x > 1 }
	fmt.Println(FilterInt([]int{0, 2}, intCheck)) // [2]
	float64Check := func(x float64) bool { return x > 0.5 }
	fmt.Println(FilterFloat64([]float64{0.0, 1.0}, float64Check)) // [1.0]
	boolCheck := func(x bool) bool { return x }
	fmt.Println(FilterBool([]bool{false, true}, boolCheck)) // [true]
	stringCheck := func(x string) bool { return len(x) > 1 }
	fmt.Println(FilterString([]string{"", "hi"}, stringCheck)) // ["hi"]
}
