package every

type Job struct {
	ctx  *Context
	unit Unit
	f    func(*Context)
}

func (j *Job) WithContext(ctx *Context) *Job {
	j.ctx = ctx
	return j
}

func (j *Job) enable(hour, minute uint8) bool {
	return j.enableHour(hour) && j.enableMinute(minute)
}

func (j *Job) enableHour(hour uint8) bool {
	if len(j.unit.hours) == 0 {
		return true
	}
	for _, h := range j.unit.hours {
		if h == hour {
			return true
		}
	}
	return false
}
func (j *Job) enableMinute(minute uint8) bool {
	if len(j.unit.minutes) == 0 {
		return true
	}
	for _, m := range j.unit.minutes {
		if m == minute {
			return true
		}
	}
	return false
}

func newJob() *Job {
	return &Job{
		ctx: NewContext(),
	}
}
