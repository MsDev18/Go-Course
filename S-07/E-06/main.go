package main

import (
	"encoding/json"
	"fmt"
)

func main() {

}

func String(err error) string {
	return fmt.Sprintln(err.Error())
}

func StringTwo(err error) string {
	return fmt.Sprintln(err)
}

type SimpleData struct {
	ID    uint
	Name  string
	Email string
}

func (s *SimpleData) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"id" : "%d", "name" : "%s" , "email" :"%s"}`, s.ID, s.Name, s.Email)), nil
}


type SimpleDataTwo struct {
	ID    uint
	Name  string
	Email string
}

func Json (data SimpleData) ([]byte,error) {
	return json.Marshal(data)
}

func JsonTwo (data SimpleDataTwo) ([]byte, error) {
	return json.Marshal(data)
}