package userservice

import (
	"E-01/dto"
	"E-01/pkg/richerror"
)

// All Request inputes for interactor/service should be sanitized
func (s Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return dto.ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
		// return ProfileResponse{}, richerror.New(err ,"userservice.Profile", err.Error(),richerror.KindUnexpected,nil)
	}
	return dto.ProfileResponse{Name: user.Name}, nil
}