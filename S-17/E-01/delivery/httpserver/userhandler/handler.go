package userhandler

import (
	"E-01/service/authservice"
	"E-01/service/userservice"
	"E-01/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
}


func New (authSvc authservice.Service, userSvc userservice.Service,userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc: authSvc,
		userSvc: userSvc,
		userValidator: userValidator,
	}
}
