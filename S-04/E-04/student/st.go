package student

import "fmt"

const STU string = "STUDENT"

func print() {
	fmt.Println("Student")
}

func AvgScore(a1, a2, a3 int) float64 {
	print()
	return float64((a1 + a2 + a3)/3)
}