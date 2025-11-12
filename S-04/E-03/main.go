package main

import (
	"fmt"
	"strings"
	"E-03/pr"
	st "E-03/pr/st"
)

func main() {
	fmt.Println("Say Hi To Go ❤️")
	s := "       Hello     "
	fmt.Println(s)
	s2 := strings.Trim(s, " ")
	fmt.Println(s2)
	pr.Print()
	var s3 Str = "fdknfgkdnfkdnfdknfdknffkn"
	s3.Print()
	st.StTest()
	var mmm Str = "fdknfgkdnfkdnfdknfdknffkn"
	mmm.Print()
	d := uint8(200)
	fmt.Println(d)
}

// Watched To S-04///E-03///1:30:00 Minutes
type Str string
func (s Str) Print() {
	fmt.Println(s)
}