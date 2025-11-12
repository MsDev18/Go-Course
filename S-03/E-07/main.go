package main

import "fmt"

func main() {
	var p *int
	i := 43

	fmt.Println("i =>", i)
	fmt.Println("p =>", p)

	p = &i
	fmt.Println("p =>", p)
	fmt.Println("value that p references to => ", *p)

	d := 73

	*p = d
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(d)
}
