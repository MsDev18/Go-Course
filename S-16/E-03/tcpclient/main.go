package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println("Start ", t)
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go connectAndSendDataToServer(i, &wg)
	}
	wg.Wait()
	fmt.Println("End ", time.Now())
	fmt.Println(time.Now().Sub(t))
	fmt.Println(time.Now())
}

func connectAndSendDataToServer(i int, wg *sync.WaitGroup) {
	defer func () {
		if r := recover() ; r != nil {
			fmt.Println("rcover : ", r)
		}
	} ()
	fmt.Println("start :", i)
	connection, err := net.Dial("tcp", ":8099")
	if err != nil {
		fmt.Println("err : ", err)
		// panic(err)
	}


	_, err = connection.Write([]byte(fmt.Sprintf(`Hello %d`, i)))
	if err != nil {
		panic(err)
	}

	connection.Close()
	wg.Done()
}
