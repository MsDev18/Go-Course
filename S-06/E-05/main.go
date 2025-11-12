package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	command := flag.String("command", "", "")
	region := flag.String("region", "", "")
	flag.Parse()
}
