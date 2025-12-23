package scheduler

import (
	"E-01/param"
	"E-01/service/matchingservice"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type Config struct {
	MatchWaitedUsersIntervalInSeconds int `koanf:"match_waited_users_interval_in_seconds"`
}

type Scheduler struct {
	config Config
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(config Config,matchSvc matchingservice.Service) Scheduler {

	return Scheduler{
		config: config,
		sch: gocron.NewScheduler(time.UTC),
		matchSvc: matchSvc,
	}
}

// Long runing process
func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	s.sch.Every(s.config.MatchWaitedUsersIntervalInSeconds).Second().Do(s.MatchWaitedUser)

	s.sch.StartAsync()
	
	<- done
	// wait to finish job 
	fmt.Println("stop scheduler ...")
	s.sch.Stop()
}

func (s Scheduler) MatchWaitedUser() {
	// matching service
	s.matchSvc.MatchWaitedUsers(param.MatchWaitedUsersRequest{})
}
