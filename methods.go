package main

import (
	"fmt"
	"math"
)

type Verte struct {
	X, Y float64
}

type CustomFloat float64

func goMethods() {
	v := Verte{2, 4}
	f := CustomFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	fmt.Println(f.Set())
	fmt.Println(formTest(&v))
}

func (v Verte) Abs() float64 {
	return math.Ceil(v.X*2 + v.X*v.Y)
}

func (f CustomFloat) Set() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func formTest(v *Verte) float64 {
	return math.Abs(v.X)
}
