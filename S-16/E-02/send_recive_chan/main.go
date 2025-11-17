package main

import (
	"fmt"
	"time"
)

// Send-Only-Chanel
func worker(done chan<- bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	reciver(done)
}

// recive-Only-Chanel
func reciver(done <-chan bool) {
	fmt.Println("Wait Start")
	d := <-done
	fmt.Println("Wait End ;  ", d)
}
