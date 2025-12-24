package matchingservice

import (
	"E-01/entity"
	"E-01/param"
	"E-01/pkg/richerror"
	"E-01/pkg/timestamp"
	"context"
	"fmt"
	"sync"
	"time"

	gofunc "github.com/thoas/go-funk"
)

// TODO - Add ctx to all repo and use-case method
type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
}

type PresenceClient interface {
	GetPresence(ctx context.Context, req param.GetPresenceRequest) (param.GetPresenceResponse, error)
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}
type Service struct {
	repo           Repository
	config         Config
	presenceClient PresenceClient
}

func New(config Config, repo Repository, presenceClient PresenceClient) Service {
	return Service{
		config:         config,
		repo:           repo,
		presenceClient: presenceClient,
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

func (s Service) MatchWaitedUsers(ctx context.Context, _ param.MatchWaitedUsersRequest) (param.MatchWaitedUsersResponse, error) {
	const op = "matchingservice.MatchWaitedUsers"

	var wg sync.WaitGroup
	for _, category := range entity.CategoryList() {
		wg.Add(1)
		go s.match(ctx, category, &wg)
	}

	wg.Wait()

	return param.MatchWaitedUsersResponse{}, nil
}

func (s Service) match(ctx context.Context, category entity.Category, wg *sync.WaitGroup) {
	const op = "matchingservice.match"

	defer wg.Done()

	list, err := s.repo.GetWaitingListByCategory(ctx, category)
	if err != nil {
		// TODO - log error
		// TODO - update metrics
		return
	}

	userIDs := make([]uint, len(list))
	for _, l := range list {
		userIDs = append(userIDs, l.UserID)
	}

	presenceList, err := s.presenceClient.GetPresence(ctx, param.GetPresenceRequest{UserIDs: userIDs})
	if err != nil {
		// TODO - log error
		// TODO - update metrics
		return
	}

	presenceUserIDs := make([]uint, len(list))
	for _, l := range presenceList.Items {
		presenceUserIDs = append(presenceUserIDs, l.UserID)
	}

	// TODO - merge presenceList with list based on userIDs
	// also consider the presence timestamp of each user
	// and remove users from waiting list if the user's timestamp is older than time.Now(-20 seconds)
	// if t < timestamp.Add(-20*time.Second) {
	// 	// remove from list
	// }	


	var finalList = make([]entity.WaitingMember, 0)
	for _, l := range list {
		if gofunc.ContainsUInt(presenceUserIDs,l.UserID) && l.Timestamp < timestamp.Add(-20*time.Second) {
			finalList = append(finalList, l)
		} else {
			// remove from waiting list 
		}
	}


	for i := 0; i < len(finalList)-1; i += 2 {
		
		mu := entity.MathcedUsers{
			Category: category,
			UserIDs:  []uint{finalList[i].UserID, finalList[i+1].UserID},
		}
		fmt.Println("mu : ", mu)
		// publishe a new event for mu
		// remove mu users from waiting list
	}

}
