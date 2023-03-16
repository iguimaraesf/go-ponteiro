package main

import "fmt"

func SomaInteiros(m map[string]int64) int64 {
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	fmt.Println(SomaInteiros(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SomaFloat(map[string]float64{"a": 1.5, "b": 2.11, "c": 3.99}))

	fmt.Println(SomaGenerica(map[string]float64{"a": 221.5, "b": 2.11, "c": 3.99}))

	var x, y, z MyNumber
	x = 4
	y = 7
	z = 31

	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
}
