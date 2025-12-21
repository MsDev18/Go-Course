package main

import (
	"E-01/config"
	"E-01/scheduler"
	"fmt"
	"os"
	"os/signal"

)

func main() {
	// TODO - read config path from command-line
	cfg := config.Load("config.yml")
	fmt.Printf("cfg : %+v\n", cfg)


	done := make(chan bool)

	go func() {
		sch := scheduler.New()
		sch.Start(done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit


	fmt.Println("recived interrupt signal, shotting down gracefully... ")
	done <- true
}
