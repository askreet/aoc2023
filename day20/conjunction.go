package day20

type Conjunction struct {
	To           []string
	LastReceived map[string]Pulse
}

func (c *Conjunction) Targets() []string {
	return c.To
}

func (c *Conjunction) AddSource(name string) {
	if c.LastReceived == nil {
		c.LastReceived = make(map[string]Pulse)
	}

	// Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
	// they initially default to remembering a low pulse for each input.
	c.LastReceived[name] = PulseLow
}

func (c *Conjunction) RecvFrom(from string, p Pulse, sendMsg func(Message)) {
	// Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
	// they initially default to remembering a low pulse for each input.
	//
	// When a pulse is received, the conjunction module first updates its memory for that input.
	c.LastReceived[from] = p

	// Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.
	for _, p := range c.LastReceived {
		if p == PulseLow {
			for _, target := range c.To {
				sendMsg(Message{Pulse: PulseHigh, To: target})
			}
			return
		}
	}

	for _, target := range c.To {
		sendMsg(Message{Pulse: PulseLow, To: target})
	}
}
