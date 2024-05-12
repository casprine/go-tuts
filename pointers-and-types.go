package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func pointersAndTypes() {
	f := finbonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// hypefunc := func(x, y float64) float64 {
	// 	return math.Sqrt(x + y)
	// }

	// // fmt.Println(hypefunc(3, 2))
	// fmt.Println(compute(hypefunc))
}

func pointers() {
	i, j := 43, 1054

	p := &i
	*p = 10

	fmt.Println(j, *p, i, p)
}

func testStruct() {
	s := Vertex{1, 2}
	r := Vertex{Y: 10}
	p := &s

	fmt.Println("X =>", s.X)
	fmt.Println(p, r)
}

func testArrays() {
	var t [2]string

	t[0] = "Jon"
	t[1] = "Appleseed"

	fmt.Println(t)
}

func testSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var rs []int = primes[2:4]
	var empt []int

	s := []struct {
		i int
		b bool
	}{
		{2, false},
		{5, true},
	}

	sr := primes[:]

	if empt == nil {
		fmt.Println("Nil")
	}

	fmt.Println(rs, s, sr, empt)
	fmt.Println(cap(rs), len(s), cap(sr))
}

func sliceWithMake() {
	rss := make([]int, 5, 10)

	fmt.Println(rss, cap(rss))
}

func ticTac() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][1] = "0"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func goingThrough() {
	names := []string{"Jon", "Desmond", "Prince", "Subject"}
	ages := []int{10, 15, 24, 245}

	fmt.Println("names =>", names)

	for i, _ := range ages {
		fmt.Printf("index => %v \n", names[i])
	}
}

func mapping() {
	type Cordinates struct {
		Lat, Long float64
	}

	var m map[string]Cordinates
	rango := map[string]Cordinates{
		"Bell": {
			24.243, 103.2422,
		},
	}
	m = make(map[string]Cordinates)

	m["Astra"] = Cordinates{
		30.5345, 24.53453,
	}

	rango["Dell"] = Cordinates{
		23.24234, 243453,
	}

	elem, ok := rango["Apple"]

	fmt.Println("Astra", m["Astra"])
	fmt.Println("Bell", rango["Bell"])
	fmt.Println("Dell", rango["Dell"])

	if ok {
		fmt.Println("I exist", elem)
	} else {
		fmt.Println("Go sit somewhere", elem)
	}

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func finbonaci() func() int {
	f2, f1 := 0, 1

	return func() int {
		f := f2
		f2, f1 = f1, f+f1

		return f
	}
}
