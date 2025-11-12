package entity

import "fmt"

type User struct {
	ID    uint
	Name  string
	Age   uint8
	Phone string
	City  string
}

func (u User) String() string {
	return fmt.Sprintf("User: {id : %d , name : %s , age : %d , phone : %s , city %s} \n", u.ID, u.Name, u.Age, u.Phone, u.City)
}
