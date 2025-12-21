package scheduler

import (
	"E-01/param"
	"E-01/service/matchingservice"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(matchSvc matchingservice.Service) Scheduler {

	return Scheduler{
		sch: gocron.NewScheduler(time.UTC),
		matchSvc: matchSvc,
	}
}

// Long runing process
func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	s.sch.Every(3).Second().Do(s.MatchWaitedUser)

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
