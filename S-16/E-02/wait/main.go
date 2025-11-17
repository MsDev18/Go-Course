package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	multipleWorkerByWaitGroup()
}

func workerByWaitGroup(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	fmt.Println("i : ", i)
	wg.Done()
}

func multipleWorkerByWaitGroup() {
	n := 5
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go workerByWaitGroup(i, &wg)
	}
	wg.Wait()
}
