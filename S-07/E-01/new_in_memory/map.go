package new_in_memory

import "E-01/user"

type Store struct {
	users map[uint]user.User
}

func (s *Store) CreateUser(u user.User) {
}

func (s *Store) ListUsers() []user.User {

}

func (s *Store) GetUserById(id uint) user.User {
	return s.users[id]
}
