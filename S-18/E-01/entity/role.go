package entity

type Role uint8

const (
	UserRole Role = iota + 1
	AdminRole
)

func (r Role) String() string {
	switch r {
	case AdminRole:
		return "Admin"
	case UserRole:
		return "user"
	}
	return ""
}
