package day20

type Untyped struct{}

func (u Untyped) RecvFrom(from string, p Pulse, sendMsg func(Message)) {
	return
}

func (u Untyped) Targets() []string {
	return []string{}
}

func (u Untyped) AddSource(name string) {
	return
}
