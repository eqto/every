package every

//Unit ..
type Unit struct {
	hours   []uint8
	minutes []uint8
}

//Minutes minutes to execute, or no param for every minute
func (u Unit) Minutes(m ...uint8) Unit {
	u.minutes = m
	if u.minutes == nil {
		u.minutes = []uint8{}
	}
	if u.hours == nil {
		u.hours = []uint8{}
	}
	return u
}

//Hours ..
func (u Unit) Hours(h ...uint8) Unit {
	u.hours = h
	if u.hours == nil {
		u.hours = []uint8{}
	}
	if u.minutes == nil {
		u.minutes = []uint8{0}
	}
	return u
}

//Run ..
func (u Unit) Run(f func(ctx Context) error) *Job {
	j := newJob()
	j.unit = u
	j.f = f
	jobLock.Lock()
	defer jobLock.Unlock()
	jobs = append(jobs, j)
	return j
}
