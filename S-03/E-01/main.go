package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  uint8
	ID   string
}

func main() {
	fmt.Println("Say Ho To Go❤️")

	users := map[int]string{
		1: "MsDev18",
		2: "Mahdieh",
		3: "Farimah",
		4: "Hasan",
	}
	fmt.Printf("Users => %v \n", users)
	v, ok := users[3]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("This Key in map Users not fount")
	}

}
