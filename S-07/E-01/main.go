package main

import (
	app "E-01/app"
	"E-01/new_in_memory"
	"E-01/user"
)

func main() {

	//var a = &storage.Memory{}

	var newMemoryStore = &new_in_memory.Store{}

	application := app.App{
		Name:      "sample-app",
		UserStore: newMemoryStore,
	}

	//u := app.User{
	//	ID:   1,
	//	Name: "MsDev18",
	//}

	application.CreateUser(user.User{
		ID:   1,
		Name: "MsDev18",
	})

}
