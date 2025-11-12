package storage

import (
	"E-01/user"
)

type Memory struct {
	users []user.User
}

func (m *Memory) CreateUser(u user.User) {
	m.users = append(m.users, u)
}

func (m *Memory) ListUsers() []user.User {
	return m.users
}

func (m *Memory) GetUserById(id uint) user.User {
	for _, user := range m.users {
		if user.ID == id {
			return user
		}
	}
	return user.User{}
}
