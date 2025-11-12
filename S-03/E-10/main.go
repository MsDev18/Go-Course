package main

import "fmt"

func main() {
	users := make([]int , 10)
	users[3] = 5
	fmt.Println("users =>",users)

	sum := 0
	for index, value := range users {
		// Dont Work 
		if index == 4 {
			value = 9
		}

		if index == 6 {
			users[index] = 8
		}
		sum += value * index
	}
	pr(users)
	fmt.Println("users => ", users)
	fmt.Println(sum)


	usersTwo := make([]int , len(users) , cap(users)) ;
	copy(usersTwo , users)
	users[8] = 80008
	fmt.Printf("users => %p | usersTwo => %p \n" , &users , &usersTwo)
	fmt.Println("value of user two => ", usersTwo)

}

func pr (s []int) {
	s[1] = 3
}

// builtin function in go 
// len() => this one for slice and array type 
// cap() => this one for slice and array type 
// copy()
// delete() => this one for map type 
// append() => this one for slice or array type 
// Watched To S-03///E-10///15:32 Minutes
// SuccesFully Watched This Epizode