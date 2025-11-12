package main

import "fmt"

func main() {
	fmt.Println("Say Hi To Go❤️")
	// prt("a","b","c","d","g")

	var slc []string = make([]string, 5, 10)
	// fmt.Println(slc)
	// slc = append(slc, "a", "b" , "c" , "MsDev18" , "Mahdieh")
	slc[0] = "a"
	slc[1] = "b"
	slc[2] = "c"
	slc[3] = "MsDev18"
	slc[4] = "Mahdieh"
	fmt.Println(slc)
	SlcPrt(slc...)
	slc = append(slc, "Mohammad", "Farimah")
	SlcPrt(slc...)
}

func SlcPrt (slc ...string) {
	// slc2 := []string{"" ,"" ,"" ,"" ,"" ,""}
	fmt.Println("slc =>", slc, "len =>", len(slc), "cap =>", cap(slc))
}
// func prt (values ...string) {
// 	fmt.Println(values)
// }
