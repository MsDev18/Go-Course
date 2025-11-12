package userservice

import (
	"E-02/entity"
	"E-02/pkg/phonenumber"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string
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

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
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
