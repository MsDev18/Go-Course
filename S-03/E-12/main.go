package main

import (
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("Hello To TODO App ❤️\n")
	command := flag.String("command", "no command", "command to run ")
	flag.Parse()

	for {
		runCommand(*command)

		fmt.Println("Please Enter another command => ")
		fmt.Scanln(command)
	}

	fmt.Println("User Storage => ", userStorage)
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}
}

func createTask() {
	var name, duedate, category string

	fmt.Printf("=> please enter the task title : ")
	fmt.Scanln(&name)

	fmt.Printf("=> please enter the task due date : ")
	fmt.Scanln(&duedate)

	fmt.Printf("=> please enter the task category : ")
	fmt.Scanln(&category)

	fmt.Printf("\nTask: name => %s ||| duedate => %s ||| category => %s\n", name, duedate, category)
}
func createCategory() {
	var title, color string

	fmt.Printf("=> please enter the category title : ")
	fmt.Scanln(&title)

	fmt.Printf("=> please enter the category color : ")
	fmt.Scanln(&color)

	fmt.Printf("Category => Title: %s ||| Color: %s  ", title, color)
}
func registerUser() {
	var id, email, password string

	fmt.Printf("=> please enter the user email : ")
	fmt.Scanln(&email)

	fmt.Printf("=> please enter the password : ")
	fmt.Scanln(&password)

	id = email

	fmt.Printf("=> User => ID: %s ||| Email: %s ||| Password => %s\n", id, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)
}
func login() {
	var id, email, password string

	fmt.Printf("=> please enter the email : ")
	fmt.Scanln(&email)

	fmt.Printf("=> please enter the password : ")
	fmt.Scanln(&password)

	fmt.Println(id, email, password)
}

// => Watched To => S-03///E-12///34:00 Minutes
