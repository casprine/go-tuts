package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	// fmt.Println(boogiewoogie())
	stackedDefer()
}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("testing index =>\n", i)
	}
}

func ifStuff(x float64) string {
	if x <= 0 {
		return fmt.Sprint(math.Pi)
	}
	return fmt.Sprint(math.Ceil((x)))
}

func boogiewoogie() {
	fmt.Print("Go runs on => ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linis")
	}
}

func stackedDefer() {
	fmt.Println("Starting the count")

	for i := 0; i < 5; i++ {
		fmt.Println("count =>", i)
		defer fmt.Println(i)
	}

	fmt.Println("Ending the count")
}
