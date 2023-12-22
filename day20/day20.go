package day20

import (
	"bufio"
	"io"
	"strings"
)

type Solution struct{}

const (
	PulseHigh = 0
	PulseLow  = 1
)

type Pulse uint8
type Messages struct {
	items []Message

	HighSendCount int
	LowSendCount  int
	InitSendCount int
}

func (m *Messages) Next() (*Message, bool) {
	if len(m.items) > 0 {
		pop := m.items[0]
		m.items = m.items[1:]
		return &pop, true
	} else {
		return nil, false
	}
}

type Message struct {
	From  string
	Pulse Pulse
	To    string
}

func (m *Messages) Send(msg Message) {
	switch msg.Pulse {
	case PulseLow:
		m.LowSendCount++
	case PulseHigh:
		m.HighSendCount++
	}
	m.items = append(m.items, msg)
}

var ButtonPress = Message{
	From:  "button",
	Pulse: PulseLow,
	To:    "broadcaster",
}

type Module interface {
	RecvFrom(from string, p Pulse, sendMsg func(Message))

	Targets() []string
	AddSource(name string)
}

type World struct {
	Modules  map[string]Module
	Messages *Messages
}

func CreateWorld(input io.Reader) World {
	var modules = make(map[string]Module)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		bits := strings.Split(line, " -> ")
		source, destConfig := bits[0], bits[1]
		switch {
		case source[0] == '%':
			list := strings.Split(destConfig, ", ")
			modules[source[1:]] = &FlipFlop{To: list}

		case source[0] == '&':
			list := strings.Split(destConfig, ", ")
			modules[source[1:]] = &Conjunction{To: list}

		case source == "broadcaster":
			list := strings.Split(destConfig, ", ")
			modules[source] = &Broadcaster{To: list}

		default:
			panic("unknown module " + source + " with config " + destConfig)
		}
	}

	w := World{
		Modules:  modules,
		Messages: &Messages{},
	}

	for name, module := range w.Modules {
		for _, target := range module.Targets() {
			if _, ok := w.Modules[target]; !ok {
				// WTF.
				w.Modules[target] = &Untyped{}
			}
			w.Modules[target].AddSource(name)
		}
	}

	return w
}

func (w *World) PressButton() {
	w.Messages.Send(ButtonPress)
}

func (w *World) Simulate() {
	for {
		msg, ok := w.Messages.Next()
		if !ok {
			break
		}

		// Provide a callback function that identifies the sender in the message,
		// since modules are not aware of their identity.
		w.Modules[msg.To].RecvFrom(msg.From, msg.Pulse, func(m Message) {
			m.From = msg.To
			w.Messages.Send(m)
		})
	}
}

func (s Solution) Part1(input io.Reader) int {
	w := CreateWorld(input)

	for i := 0; i < 1000; i++ {
		w.PressButton()
		w.Simulate()
	}

	return w.Messages.LowSendCount * w.Messages.HighSendCount
}

func (s Solution) Part2(input io.Reader) int {
	w := CreateWorld(input)

	rx := &CountLows{}
	w.Modules["rx"] = rx

	ct := 0
	for {
		ct++
		w.PressButton()
		w.Simulate()
		if rx.LowCount > 0 {
			return ct
		}
	}
}
