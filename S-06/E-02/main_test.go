package main

import "testing"

//func TestSquare(t *testing.T) {
//	for i := 1; i < 10000; i++ {
//		result := square(i)
//		if result/i != i {
//			t.Errorf("exepted : %d, result : %d \n", i*i, result)
//		}
//	}
//}

func TestDayOfWeek(t *testing.T) {
	type test struct {
		Input          int
		ExpectedResult string
	}
	var testCases = []test{
		{Input: 1, ExpectedResult: "شنبه"},
		{Input: 2, ExpectedResult: "یک شنبه"},
		{Input: 3, ExpectedResult: "دو شنبه"},
		{Input: 4, ExpectedResult: "سه شنبه"},
		{Input: 5, ExpectedResult: "چهار شنبه"},
		{Input: 6, ExpectedResult: "پنج شنبه"},
		{Input: 7, ExpectedResult: "جمعه"},
		{Input: 8, ExpectedResult: ""},
		{Input: 0, ExpectedResult: ""},
		{Input: -1, ExpectedResult: ""},
	}

	for _, c := range testCases {
		result := dayOfWeek(c.Input)
		if result != c.ExpectedResult {
			t.Errorf("exepted : %s, result : %s \n", c.ExpectedResult, result)
		}
	}
}
