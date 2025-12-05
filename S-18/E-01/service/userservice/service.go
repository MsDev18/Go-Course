package userservice

import (
	"E-01/entity"
	"crypto/md5"
	"encoding/hex"
)

type Repository interface {

	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{
		repo: repo,
		auth: authGenerator,
	}
}


func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

