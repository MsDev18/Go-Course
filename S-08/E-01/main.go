package main

import "fmt"

func main() {
	fmt.Println("-----<<< Say Hi To Go 😍 >>>-----")

	u := Teacher{
		ID:       1,
		Name:     "Msdev18",
		Email:    "MsDev18@gmail.com",
		IsActive: true,
	}

	fmt.Println("User : ", u)
}

type Teacher struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
	Salary   float64
	Tax      float64
}

type Student struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

type DeActiveUser interface{
	DeActiveUser() error
}

func (t *Teacher) DeActiveUser () {
	if t.IsActive {
		
	}
}

func (s *Student) DeActiveUser () {
	
}