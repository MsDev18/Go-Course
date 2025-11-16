package main

import (
	"fmt"
	"time"
)

// import "fmt"

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func sum2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}


	// sender
	c <- sum
	c <- sum

	fmt.Println("End sum2 : ", sum)
}

func main() {
	var c chan int = make(chan int, 1)

	s := []int{4, -9, 25, 34, -15, 14}

	// go sum2(s[:len(s)/2],c)

	go sum2(s[len(s)/2:],c)

	// reciver
	sum := <- c
	fmt.Println("sum : ", sum)
	time.Sleep(time.Second * 1)

	// sum = <- c
	// fmt.Println("sum : ",sum)

	// sum = <- c
	// fmt.Println("sum : ",sum)
}
