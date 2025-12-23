package matchinghandler

import (
	"E-01/service/authservice"
	"E-01/service/matchingservice"
	"E-01/service/presenceservice"
	"E-01/validator/matchingvalidator"
)

type Handler struct {
	authConfig        authservice.Config
	authSvc           authservice.Service
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	presenceSvc       presenceservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator, presenceSvc presenceservice.Service) Handler {
	return Handler{
		authConfig:        authConfig,
		authSvc:           authSvc,
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
		presenceSvc:       presenceSvc,
	}
}
