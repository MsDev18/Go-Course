package scheduler

import (
	"E-01/param"
	"E-01/service/matchingservice"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type Config struct {
	MatchWaitedUsersIntervalInSeconds int `koanf:"match_waited_users_interval_in_seconds"`
}

type Scheduler struct {
	config   Config
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(config Config, matchSvc matchingservice.Service) Scheduler {

	return Scheduler{
		config:   config,
		sch:      gocron.NewScheduler(time.UTC),
		matchSvc: matchSvc,
	}
}

// Long runing process
func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	s.sch.Every(s.config.MatchWaitedUsersIntervalInSeconds).Second().Do(s.MatchWaitedUser)

	s.sch.StartAsync()

	<-done
	// wait to finish job
	fmt.Println("stop scheduler ...")
	s.sch.Stop()
}

func (s Scheduler) MatchWaitedUser() {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	// get lock
	_, err := s.matchSvc.MatchWaitedUsers(ctx, param.MatchWaitedUsersRequest{})
	if err != nil {
		// TODO - log err
		// TODO - update metrics 
		fmt.Println("matchSvc.MatchWaitedUsers error : ", err) 
	}
	// free lock
}
