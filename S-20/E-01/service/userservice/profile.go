package userservice

import (
	"E-01/param"
	"E-01/pkg/richerror"
	"context"
)

// All Request inputes for interactor/service should be sanitized
func (s Service) Profile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}
	return param.ProfileResponse{Name: user.Name}, nil
}