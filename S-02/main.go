package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Say Hi To Go ❤️")

	firstName := flag.String("word", "default value", "description")
	age := flag.Int("age", 23, "age of user")

	flag.Parse()

	fmt.Printf("User name is => %s \n", *firstName)
	fmt.Printf("User age is => %d \n", *age)

	rFlags := flag.Args()
	fmt.Printf("len remaining flags %v \n", len(rFlags))
	fmt.Printf("len remaining flags %v \n", rFlags)

	// fmt.Println("\nPlease enter your name: ")
	// var first_name, last_name string
	// fmt.Scanln(&first_name, &last_name)
	// fmt.Println("first name => ", first_name, last_name)

	fmt.Printf("1 , 2, 4 , 6 , 8 \n")
	var a1, a2, a3, a4, a5 int
	fmt.Scanf("%d,%d,%d,%d,%d", &a1, &a2, &a3, &a4, &a5)
	fmt.Println(a1, a2, a3, a4, a5)
}
