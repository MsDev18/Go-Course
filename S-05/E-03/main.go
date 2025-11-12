package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDon      bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

// Global
var userStorage []User
var authenticatedUser *User
var taskStorage []Task
var categoryStorage []Category

func main() {
	fmt.Println("Hello To TODO App ❤️")

	command := flag.String("command", "no command", "command to run ")
	flag.Parse()
	fmt.Println(*command)

	// if there is a user record with corresponding data allow the user to continue

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please Enter another command => ")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, duedate, category string

	fmt.Println("please enter the task title : ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("=> please enter the task category id : ")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("category-id is not valid integer", err)
		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Println("category-id is not found ")

		return
	}

	fmt.Println("=> please enter the task due date : ")
	scanner.Scan()
	duedate = scanner.Text()

	// validation
	// category validated

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDon:      false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("=> please enter the category title : ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("=> please enter the category color : ")
	scanner.Scan()
	color = scanner.Text()

	c := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Printf("=> please enter the user name : ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("=> please enter the user email : ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Printf("=> please enter the password : ")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("User : ", id, name, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)
}

func login() {
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Printf("=> please enter the email : ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Printf("=> please enter the password : ")
	scanner.Scan()
	password = scanner.Text()

	// Get The Email And Password From The Client
	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("The email or password is not correct")
	}
}

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
