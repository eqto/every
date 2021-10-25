package every

type Unit struct {
	hours   []uint8
	minutes []uint8
}

func (u Unit) Do(f func(*Context)) *Job {
	j := newJob()
	j.unit = u
	j.f = f
	jobLock.Lock()
	defer jobLock.Unlock()
	jobs = append(jobs, j)
	run()
	return j
}
