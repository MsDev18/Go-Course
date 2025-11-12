package app

import (
	"E-01/user"
	"fmt"
)

type userStore interface {
	CreateUser(u user.User)
	ListUsers() []user.User
	GetUserById(id uint) user.User
}

type App struct {
	Name      string
	UserStore userStore
}

func (a App) CreateUser(u user.User) {
	if u.Name == "" {
		fmt.Println("name can't be empty")

		return
	}
	//var fileHandler *os.File
	//
	//if f, err := os.OpenFile(a.StorageFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err != nil {
	//	fmt.Println("can't open the file : ", err)
	//
	//	return
	//} else {
	//	fileHandler = f
	//}
	//defer func() {
	//	dErr := fileHandler.Close()
	//	if dErr != nil {
	//		fmt.Println("can't close the file : ", dErr)
	//	}
	//}()

	// serialize user
	//data, mErr := json.Marshal(u)
	//if mErr != nil {
	//	fmt.Printf("can't marshal user data %v \n", mErr)
	//
	//	return
	//}

	//a.InMemoryStorage.AddUser(u)

	a.UserStore.CreateUser(u)

	//if _, wErr := fileHandler.Write(data); wErr != nil {
	//	fmt.Println("can't write to the file : ", wErr)
	//
	//return
	//}

}
