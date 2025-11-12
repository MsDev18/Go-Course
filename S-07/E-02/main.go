package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func main() {
	name := "MsDev18"
	stringReader := strings.NewReader(name)
	scanner := bufio.NewScanner(stringReader)

	scanner.Scan()
	fmt.Println("output : ")
	fmt.Println(scanner.Text())

	var scores = Int{12, 1, 3, 4, 25, 12, 143, 10, 99, 50, 40}
	fmt.Println("before sort : ", scores)
	sort.Sort(scores)
	fmt.Println("after sort : ", scores)
}

type Int []int

func (in Int) Len() int {
	return len(in)
}
func (in Int) Less(i, j int) bool {
	return in[i] < in[j]
}
func (in Int) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

type User struct {
	ID   uint
	Name string
}
type userStore map[uint]User
