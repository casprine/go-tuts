package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	// fmt.Println(inferedVariables())
	typeConversion()
}

func GetVersion() {
	fmt.Println("This current version is", rand.Intn(10))
}

func add(x, y int) int {
	return x + y
}

func namedReturns(fc int) (by2, by4 int) {
	by2 = fc * 2
	by4 = fc * 4

	return
}

func variables() (bool, bool, bool) {
	var c, python, java bool

	return c, python, java
}

func inferedVariables() (int, bool) {
	var count = 43
	var isOpened = false

	return count, isOpened
}

func tryingTypes() {
	var (
		isBool           bool       = false
		MaxInt           uint       = 1<<64 - 1
		somethingComplex complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type %T => Value: %v\n", isBool, isBool)
	fmt.Printf("Type %T => Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T => Value: %v\n", somethingComplex, somethingComplex)
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
