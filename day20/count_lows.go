package day20

type CountLows struct {
	LowCount int
}

func (c CountLows) RecvFrom(from string, p Pulse, sendMsg func(Message)) {
	if p == PulseLow {
		c.LowCount++
	}
}

func (c CountLows) Targets() []string {
	return []string{}
}

func (c CountLows) AddSource(name string) {
	return
}
