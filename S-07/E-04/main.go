package main

import (
	"E-04/log"
	"E-04/richerror"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID   uint
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{ id: %d , name : %s } \n", u.ID, u.Name)
}

func main() {
	//u := User{
	//	ID:   123,
	//	Name: "MsDev18",
	//}
	//
	//fmt.Println(u)

	logger := log.Log{}

	_, oErr := os.OpenFile("./storage/data.txt", os.O_RDWR, 0666)
	if oErr != nil {
		// logger.Append(oErr)

		fmt.Println(oErr.Error())
	}

	_, gErr := getUserByID(0)
	if gErr != nil {
		// type assertion
		rErr, ok := gErr.(*richerror.RichError)
		if ok {
			logger.Append(rErr)
		} else {
			logger.Append(&richerror.RichError{
				Message:   gErr.Error(),
				MetaData:  nil,
				Operation: "unknown",
			})
		}
	}

	_, g2Err := getUserByIDTwo(0)
	fmt.Println("Operation : ", g2Err.Operation)
	if g2Err != nil {
		logger.Append(g2Err)
	}

	_, g3Err := getUserByIDThree(0)
	_,ok := g3Err.(SimpleError)
	if !ok {
		
	}

	logger.Save()
}

// abstraction type
func getUserByID(id int) (User, error) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByID",
		}
	}
	return User{}, nil
}

// concrete type
func getUserByIDTwo(id int) (User, *richerror.RichError) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByIDTwo",
		}
	}
	return User{}, nil
}

type SimpleError struct {
	Output    string
	Operation string
}


func (s SimpleError) Error () string {
	return "output : " + s.Output + " , operation : " + s.Operation 
}

func getUserByIDThree(id int) (User, error) {
	if id == 0 {
		return User{}, &SimpleError{
			Output: "id is 0",
			Operation: "getUserByIdThree",
		}
	}
	return User{}, nil
}