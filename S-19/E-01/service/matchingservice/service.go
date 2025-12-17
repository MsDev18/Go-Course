package matchingservice

import (
	"E-01/entity"
	"E-01/param"
	"E-01/pkg/richerror"
	"time"
)

type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}
type Service struct {
	repo   Repository
	config Config
}

func New(config Config,repo Repository) Service {
	return Service{
		config: config,
		repo: repo,
	}
}

func (s Service) AddToWatingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitingListList"

	// add user to waiting ;list for the given category if not exixt
	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return param.AddToWaitingListResponse{
		Timeout: s.config.WaitingTimeout,
	}, nil
}
