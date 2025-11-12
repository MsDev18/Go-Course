package main

import "fmt"

var (
	reset string = "\u001B[0m"
	red   string = "\u001B[31m"
	blue  string = "\033[34m"
)

func main() {
	fmt.Printf("%s Say Hi To Go ... %s \n", blue, reset)
}
