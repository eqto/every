package every

import (
	"sync"
	"time"
)

var (
	tickCallback func(time.Time)
	jobs         []*Job
	jobLock      = sync.Mutex{}
	wait         = make(chan struct{}, 1)
	jobWait      = sync.WaitGroup{}

	done     chan int
	doneLock = sync.Mutex{}
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

//Wait ..
func Wait() {
	<-wait
}

func run() {
	doneLock.Lock()
	defer doneLock.Unlock()
	if done != nil {
		return
	}

	done = make(chan int, 1)
	go func() {
		for done != nil {
			runFunc(done)
		}
		jobWait.Wait()
		wait <- struct{}{}
	}()
}

//Stop ..
func Stop() {
	doneLock.Lock()
	defer doneLock.Unlock()
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

		jobWait.Add(len(jobs))
		for _, job := range jobs {
			go func(job *Job) {
				defer jobWait.Done()
				if job.enable(hour, minute) {
					job.f(job.ctx)
				}
			}(job)
		}
	case <-cancel:
		doneLock.Lock()
		defer doneLock.Unlock()
		done = nil
	}
	if tickCallback != nil {
		tickCallback(next)
	}
}
