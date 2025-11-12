package main

import "fmt"

func main() {
	//a := 42
	//pr(a)
	//fmt.Println("aaaa", a)
	//
	//prP(&a)
	//fmt.Println("2aaa", a)

	s := Student{
		Name: "MsDev18",
	}
	fmt.Println(s.getName())
	s.setName("test")
	fmt.Println(s.getName())
	fmt.Println(s)
}

func pr(a int) {
	a += 10
	fmt.Println(a)
}
func prP(a *int) {
	*a += 22
	fmt.Println(*a)
}

type Student struct {
	Name string
}

//	func (s Student) print() {
//		fmt.Println(s.Name)
//	}
func (s *Student) print() {
	fmt.Println(s.Name)
}

// Geter Or Setter
func (s Student) getName() string {
	return s.Name
}
func (s *Student) setName(name string) {
	s.Name = name
}

// Watched S-03///E-09///11:00 Minutes
