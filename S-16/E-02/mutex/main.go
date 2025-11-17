package main

import (
	"fmt"
	"sync"
)

func write(m map[string]int, key string, value int, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	m[key] += value
	mx.Unlock()
}

func main() {
	m := make(map[string]int)
	n := 10
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go write(m, "counter", i, &wg, &mx)
	}
	wg.Wait()

	fmt.Println(m["counter"])
}
