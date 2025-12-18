package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
	jobs []string
}

func New() Scheduler {
	return Scheduler{}
}

// Long runing process
func (s Scheduler) Start(done <- chan bool) {
	fmt.Println("Scheduler Start")
	for {
		select {
		case d := <-done:
			fmt.Println("Exiting ...", d)
			return
		default:
			now := time.Now()
			fmt.Println("Scheduler Now :=> ", now)
			time.Sleep(time.Second * 3)
		}
	}
}
