package main

import (
	"E-05/entity"
	"fmt"
)

func main() {
	fmt.Println("Hello World ...")
	u := entity.User{
		ID: 1,
		Name: "MsDev18",
		Age: 18,
		Phone: "09351721415",
		City: "Tehran",
	}
	fmt.Println(u)
	
}
