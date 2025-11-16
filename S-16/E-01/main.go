package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine
	go func1()
	go func2()


	time.Sleep(time.Second * 5)
}

func func1() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Func 1 -> i => " , i)
	}
}

func func2() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Func 2 -> i => " , i)
	}
}