package main

import (
	"bufio"
	"encoding/json"
	"errors"
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
var (
	userStorage       []User
	authenticatedUser *User
	taskStorage       []Task
	categoryStorage   []Category
	serializationMode string
)

const (
	USER_STORAGE_PATH                 string = "./user.txt"
	MAN_DAR_AVARDI_SERIALIZATION_MODE string = "mandaravardi"
	JSON_SERIALIZATION_MODE           string = "json"
)

func main() {
	serializaMode := flag.String("serialize-mode", MAN_DAR_AVARDI_SERIALIZATION_MODE, "serialization mode to write data file")
	command := flag.String("command", "no command", "command to run ")
	flag.Parse()
	// load user storage from file
	loadUserStorageFromFile(serializationMode)

	fmt.Println("Hello To TODO App ❤️")

	switch *serializaMode {
	case MAN_DAR_AVARDI_SERIALIZATION_MODE:
		serializationMode = MAN_DAR_AVARDI_SERIALIZATION_MODE
	default:
		serializationMode = JSON_SERIALIZATION_MODE
	}

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
	writeUserToFile(user)
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

func loadUserStorageFromFile(serializationMode string) {
	file, err := os.OpenFile(USER_STORAGE_PATH, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)

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

	for _, u := range userSlice {
		var userStruct = User{}

		switch serializationMode {
		case MAN_DAR_AVARDI_SERIALIZATION_MODE:
			var dErr error
			userStruct, dErr = deserializeFromManDarAvardi(u)
			if dErr != nil {
				fmt.Println("can't deserialize user record to user struct ", dErr)

				return
			}
		case JSON_SERIALIZATION_MODE:
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("can't deserialize user record to user struct with json mode", uErr)

				return
			}
		default:
			fmt.Println("Error ")
			return
		}
		fmt.Println(userStruct)
		userStorage = append(userStorage, userStruct)
	}

	// fmt.Println(data)
}

func writeUserToFile(user User) {

	file, err := os.OpenFile(USER_STORAGE_PATH, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println("error in open file with os ", err)

		return
	}

	defer file.Close()

	// Serialize the user struct || object
	var data []byte
	if serializationMode == MAN_DAR_AVARDI_SERIALIZATION_MODE {
		data = []byte(fmt.Sprintf("id : %d, name : %s, email : %s, password : %s \n", user.ID, user.Name, user.Email, user.Password))
	} else if serializationMode == JSON_SERIALIZATION_MODE {
		//json
		var jErr error
		data, jErr = json.Marshal(user)
		if jErr != nil {
			fmt.Println("can't marshal user struct to json", err)

			return
		}
	} else {
		fmt.Println("inValid Serialization mode")

		return
	}

	_, wErr := file.Write(data)
	if wErr != nil {
		fmt.Printf("can't write to the file %v \n", wErr)
	}

}
func deserializeFromManDarAvardi(userStr string) (User, error) {
	var user User
	if userStr == "" {
		return User{}, errors.New("user string id empty")
	}
	fmt.Println("user", userStr)

	userFields := strings.Split(userStr, ",")

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
				return User{}, errors.New("strconv err")
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
	fmt.Printf("user : %+v", user)
	return user, nil
}
