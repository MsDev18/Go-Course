package userservice

import (
	"E-01/dto"
	"E-01/entity"
	"fmt"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    getMD5Hash(req.Password),
	}
	// Create New User In Storage
	// TODO - assign anonymose struct
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
	}
	// Return Created User

	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}