package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	for i := 0; i <= 100; i++ {

		go F1(i)
	}
	fmt.Println("runtime.NumGoroutine() 1 : ", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
	fmt.Println("runtime.NumGoroutine() 2 : ", runtime.NumGoroutine())
	runtime.Gosched()
}

func F1(i int) {
	fmt.Println("Helllo : ", i)
}
