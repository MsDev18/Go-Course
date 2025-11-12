package main

import (
	"fmt"
)

func main() {
	fmt.Println("Say Hi To Go ❤️")

	s := make([]int, 3, 12)
	
	s = append(s, []int{1, 2, 3}...)
	fmt.Println(s)
	s = append(s, []int{5,6,9}...)
	fmt.Println(s)
}
