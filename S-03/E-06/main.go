package main

import (
	"fmt"
)

func main() {
	fmt.Println("Say Hi To Go ❤️")

	name1 := "MsDev18"
	name2 := &name1
	fmt.Println("name 1 => ", name1)
	fmt.Println("name 2 => ", *name2)
	name1 = "Mohammad"
	fmt.Println("name 1 => ", name1)
	fmt.Println("name 2 => ", *name2)
}

// Finished Watch
