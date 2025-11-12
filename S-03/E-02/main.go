package main

import (
	"fmt"
	"time"
)

func print(d int) {
	fmt.Println("d", d)
}

func t(d int) {
	fmt.Println("time", time.Now(), d)
}

func main() {
	// Function Values
	f := func(b, p int) int {
		return b + p
	}

	fmt.Println(f(2, 4))

	compute(6, print)
	compute(8, t)
	compute(12, func(k int) {
		fmt.Println("k", k*2)
	})

	func (a int) {
		fmt.Println("Test for anonymouse function", a)
	} (70)
}

func compute(b int, f func(r int)) (result int) {
	result = b * 2
	f(result)
	return
}
