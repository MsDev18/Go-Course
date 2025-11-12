package main

import "fmt"

func main() {
	fmt.Println("Say Hi To Go ❤️")
	// Method

	// Instance of Student Type
	st := Student{
		Name:   "MsDev18",
		ID:     "0151509298",
		Scores: []int{5, 12, 11, 9, 16},
	}
	// functions run
	st.printStudent()
	fmt.Println(st.studentAvgScore())
	fmt.Println(st.isStudentEligible())
}

type Student struct {
	Name   string
	ID     string
	Scores []int
}

/// Declaration Functionality
/////////////////////////////
/////////////////////////////
/////////////////////////////
//func printStudent(s Student) {
//	fmt.Printf("The student name is %s and the id is %s \n", s.Name, s.ID)
//}
//
//func studentAvgScore(s Student) (avg float64) {
//	sum := 0
//	for _, s := range s.Scores {
//		sum += s
//	}
//	avg = float64(sum) / float64(len(s.Scores))
//	return
//}
//
// func isStudentEligible(s Student) bool {
//	if studentAvgScore(s) > 12 {
//		return true
//	}
//
//	return false
//}

// Declaration With Method
// ////////////////////////
// ////////////////////////
// ////////////////////////
func (s Student) printStudent() {
	fmt.Printf("The student name is %s and the id is %s \n", s.Name, s.ID)
}
func (s Student) studentAvgScore() (avg float64) {
	sum := 0
	for _, s := range s.Scores {
		sum += s
	}
	avg = float64(sum) / float64(len(s.Scores))
	return
}
func (s Student) isStudentEligible() bool {
	if s.studentAvgScore() > 12 {
		return true
	}
	return false
}

// Watched Finished
