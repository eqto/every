package every

import (
	"sync"
	"time"
)

var (
	tickCallback func(time.Time)
	cancelFunc   func(*Context) bool
	jobs         []*Job
	jobLock      = sync.Mutex{}
	wait         = make(chan struct{}, 1)
	jobWait      = sync.WaitGroup{}

	done     chan struct{}
	doneLock = sync.Mutex{}
)

func TickCallback(f func(time.Time)) {
	tickCallback = f
}

// Hours hours to execute, or no param for every hours
func Hours(hours ...uint8) Hour {
	hour := Hour{}
	if len(hours) > 0 {
		hour.unit.hours = hours
	}
	hour.unit.minutes = []uint8{0}
	return hour
}

// Minutes minutes to execute, or no param for every minutes
func Minutes(minutes ...uint8) Minute {
	minute := Minute{}
	if len(minutes) > 0 {
		minute.unit.minutes = minutes
	}
	return minute
}

// cancel will executed before each job run and stop execution when cancel return true
func Precheck(cancel func(ctx *Context) bool) {
	cancelFunc = cancel
}

// Wait wait for any jobs still running to finish
func Wait() {
	<-wait
}

func run() {
	doneLock.Lock()
	defer doneLock.Unlock()
	if done != nil {
		return
	}

	done = make(chan struct{}, 1)
	go func() {
		for done != nil {
			runFunc(done)
		}
		jobWait.Wait()
		wait <- struct{}{}
	}()
}

func Stop() {
	doneLock.Lock()
	defer doneLock.Unlock()
	if done == nil {
		return
	}
	done <- struct{}{}
}

func runFunc(cancel <-chan struct{}) {
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
					if cancelFunc == nil || (cancelFunc != nil && !cancelFunc(job.ctx)) {
						job.ctx.hour = hour
						job.ctx.minute = minute
						job.f(job.ctx)
					}
				}
			}(job)
		}
	case <-cancel:
		doneLock.Lock()
		defer doneLock.Unlock()
		done = nil
	}
	if tickCallback != nil {
		go tickCallback(next)
	}
}
