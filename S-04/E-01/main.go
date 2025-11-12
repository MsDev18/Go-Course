package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Println("Say Ho To GoLang ... ❤️")

	a := 25235
	var f int = 122
	fmt.Printf("%T \n", a)
	fmt.Printf("%d \n", unsafe.Sizeof(a))

	fmt.Println(runtime.GOOS, runtime.GOARCH)

	fmt.Println("Hello", f)
	var s []rune = []rune("سلام دنیا")
	fmt.Println(len("سلام دنیا"))
	fmt.Println(s)
}

// Watched SuccessFully
// int8 => 1
// int16 => 2
// int32 => 4
// int64 => 8
