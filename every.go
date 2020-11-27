package every

import (
	"sync"
	"time"

	"github.com/eqto/command"
)

var (
	tickCallback func(time.Time)
	jobs         []*Job
	jobLock      = sync.Mutex{}
)

//TickCallback ..
func TickCallback(f func(time.Time)) {
	tickCallback = f
}

//Minutes minutes to execute, or no param for every minute
func Minutes(m ...uint8) Unit {
	return Unit{}.Minutes(m...)
}

//Hours ..
func Hours(h ...uint8) Unit {
	return Unit{}.Hours(h...)
}

//Start ..
func Start() error {
	command.Add(runFunc, 1)
	if e := command.Start(); e != nil {
		return e
	}
	command.Wait()
	return nil
}

func runFunc(done <-chan int) {
	next := time.Unix(time.Now().Unix(), 0)
	next = next.Add(time.Duration(60-next.Second()) * time.Second)

	select {
	case t := <-time.After(time.Until(next)):
		jobLock.Lock()
		defer jobLock.Unlock()

		hour := uint8(t.Hour())
		minute := uint8(t.Minute())

		for _, job := range jobs {
			go func(job *Job) {
				if job.enable(hour, minute) {
					job.f(job.ctx)
				}
			}(job)
		}
	case <-done:
	}
	if tickCallback != nil {
		tickCallback(next)
	}
}
