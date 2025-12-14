package userservice

import (
	"E-01/entity"
	"E-01/param"
	"fmt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    getMD5Hash(req.Password),
		Role:        entity.UserRole,
	}
	// Create New User In Storage
	// TODO - assign anonymose struct
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
	}
	// Return Created User

	return param.RegisterResponse{User: param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}
