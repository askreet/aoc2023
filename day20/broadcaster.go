package day20

type Broadcaster struct {
	To []string
}

func (b *Broadcaster) Targets() []string {
	return b.To
}

func (b *Broadcaster) AddSource(name string) {
	return
}

func (b *Broadcaster) RecvFrom(from string, p Pulse, sendMsg func(Message)) {
	for _, target := range b.To {
		sendMsg(Message{To: target, Pulse: p})
	}
}
