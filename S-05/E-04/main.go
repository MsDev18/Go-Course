package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

const USER_STORAGE_PATH string = "./user.txt"

func main() {

	// load user storage from file
	loadUserStorageFromFile()
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
	// save user data in user .txt
	var file *os.File

	file, err := os.OpenFile(USER_STORAGE_PATH, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println("error in open file with os ", err)

		return
	}

	data := fmt.Sprintf("id : %d, name : %s, email : %s, password : %s \n", user.ID, user.Name, user.Email, user.Password)

	var b = []byte(data)
	numberOfWrittenBytes, wErr := file.Write(b)
	if wErr != nil {
		fmt.Printf("can't write to the file %v \n", wErr)
	}

	fmt.Println("numberOfWrittenBytes : ", numberOfWrittenBytes)
	file.Close()
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

func loadUserStorageFromFile() {
	file, err := os.OpenFile(USER_STORAGE_PATH, os.O_RDONLY, 0777)

	if err != nil {
		fmt.Println("can't open the file ", err)
	}

	var data []byte = make([]byte, 1024)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file ", oErr)
	}
	var dataStr = string(data)

	userSlice := strings.Split(dataStr, "\n")
		var user = User{}

	for index, u := range userSlice {
		if u == "" {
			continue
		}
		fmt.Println("line of file ", index, "user", u)

		userFields := strings.Split(u, ",")

		for _, field := range userFields {
			fmt.Println(field)
			values := strings.Split(field, " : ")
			if len(values) != 2 {
				fmt.Println("field is not valid, skipping...")
				continue
			}
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]
			fmt.Println(fieldName, fieldValue)

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					fmt.Println("strconv err ", err)

					return
				}
				user.ID = id
			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue
			}
		}
	}
	fmt.Printf("user : %+v", user)

	// fmt.Println(data)
}
