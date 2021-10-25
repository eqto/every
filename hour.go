package every

type Hour struct {
	unit Unit
}

func (h Hour) Minutes(minutes ...uint8) Minute {
	m := Minutes(minutes...)
	m.unit.hours = h.unit.hours
	return m
}

func (h Hour) Excepts(hours ...uint8) Hour {
	if len(h.unit.hours) == 0 {
		for i := 0; i < 24; i++ {
			h.unit.hours = append(h.unit.hours, uint8(i))
		}
	}
	m := make(map[uint8]struct{})
	for _, hour := range hours {
		m[hour] = struct{}{}
	}
	new := []uint8{}
	for _, hour := range h.unit.hours {
		if _, ok := m[hour]; !ok {
			new = append(new, hour)
		}
	}
	h.unit.hours = new
	return h
}

func (h Hour) Do(f func(*Context)) *Job {
	return h.unit.Do(f)
}
