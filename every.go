package every

import (
	"sync"
	"time"
)

var (
	tickCallback func(time.Time)
	jobs         []*Job
	jobLock      = sync.Mutex{}
	runLock      = sync.Mutex{}
	done         chan int
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

func run() {
	runLock.Lock()
	defer runLock.Unlock()

	if done != nil {
		return
	}
	done = make(chan int, 1)
	go func() {
		for done != nil {
			runFunc(done)
		}
	}()
}

func stop() {
	runLock.Lock()
	defer runLock.Unlock()
	if done == nil {
		return
	}
	done <- 1
}

func runFunc(cancel <-chan int) {
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
	case <-cancel:
		runLock.Lock()
		defer runLock.Unlock()
		done = nil
	}
	if tickCallback != nil {
		tickCallback(next)
	}
}
