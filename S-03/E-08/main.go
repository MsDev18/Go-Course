package main

import (
	"E-08/student"
	"fmt"
)

func main() {
	fmt.Println("Say Hi To Go❤️")
	var St student.Student = student.Student{
		Contry: "Iran",
		City:   "Tehran",
		Age:    18,
		Name:   "MsDev18",
	}
	St.Prt()
}
