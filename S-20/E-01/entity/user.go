package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// Password Always Keep Hashed Password
	Password string
	Role Role
}
