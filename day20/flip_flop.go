package day20

type FlipFlop struct {
	To []string
	On bool
}

func (f *FlipFlop) Targets() []string {
	return f.To
}

func (f *FlipFlop) AddSource(name string) {
	return
}

func (f *FlipFlop) RecvFrom(from string, p Pulse, sendMsg func(Message)) {
	// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
	if p == PulseHigh {
		return
	}

	// However, if a flip-flop module receives a low pulse, it flips between on and off.
	if !f.On {
		// If it was off, it turns on and sends a high pulse.
		f.On = true
		for _, target := range f.To {
			sendMsg(Message{Pulse: PulseHigh, To: target})
		}
	} else {
		// If it was on, it turns off and sends a low pulse.
		f.On = false
		for _, target := range f.To {
			sendMsg(Message{Pulse: PulseLow, To: target})
		}
	}
}
