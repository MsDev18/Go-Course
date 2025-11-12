package userservice

import (
	"E-01/entity"
	"E-01/pkg/phonenumber"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

	// TODO - we should verify phone number with verification code
	// Validate Phone Number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}
	// Check Uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("number is not unique")
		}
	}
	// Validate Name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be greatest than 3")
	}

	// TODO - Check The Password With Regex Pattern
	// Validate Password
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password len should be greater than 8")
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    getMD5Hash(req.Password),
	}
	// Create New User In Storage
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
	}
	// Return Created User
	return RegisterResponse{
		User: createdUser,
	}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user two separate method for existence check and getUserByPhoneNumber
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	// 3. Compare user.password with the req.password
	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	// 4. return ok
	return LoginResponse{}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
