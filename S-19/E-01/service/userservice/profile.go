package userservice

import (
	"E-01/param"
	"E-01/pkg/richerror"
)

// All Request inputes for interactor/service should be sanitized
func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
		// return ProfileResponse{}, richerror.New(err ,"userservice.Profile", err.Error(),richerror.KindUnexpected,nil)
	}
	return param.ProfileResponse{Name: user.Name}, nil
}