package every

type Minute struct {
	unit Unit
}

func (m Minute) Excepts(minutes ...uint8) Minute {
	if len(m.unit.minutes) == 0 {
		for i := 0; i < 60; i++ {
			m.unit.minutes = append(m.unit.minutes, uint8(i))
		}
	}
	minuteMap := make(map[uint8]struct{})
	for _, minute := range minutes {
		minuteMap[minute] = struct{}{}
	}
	new := []uint8{}
	for _, minute := range m.unit.minutes {
		if _, ok := minuteMap[minute]; !ok {
			new = append(new, minute)
		}
	}
	m.unit.minutes = new
	return m
}

func (m Minute) Do(f func(*Context)) *Job {
	return m.unit.Do(f)
}
