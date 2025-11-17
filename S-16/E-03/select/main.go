package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	done := make(chan bool)
	err := make(chan error)
}

func worker(done chan<- bool, err chan<- error) {
	n := rand.Int(10)
	if n < 5 {
		done <- true
	} else {
		err <- fmt.Errorf("value is greater than 5")
	}
}