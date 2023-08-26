package main

import (
	"fmt"
	"math"
)

type Machine struct {
	tape           []rune
	head           int
	curr           rune
	configurations map[rune]Configuration
	halt           bool
	firstType      map[rune]bool // I use a bool because we don't have native hashsets
}

func NewMachine(configurations map[rune]Configuration, startingConfiguration rune, firstTypeSymbols map[rune]bool) *Machine {
	tapeSize := int(math.Pow(2, 4)) - 1
	tape := make([]rune, tapeSize)
	for i := 0; i < tapeSize; i += 1 {
		tape[i] = None
	}

	return &Machine{
		tape:           tape,
		head:           int(math.Pow(2, 3)),
		curr:           startingConfiguration,
		configurations: configurations,
		halt:           false,
		firstType:      firstTypeSymbols,
	}
}

func (m *Machine) step() {
	if m.halt {
		return
	}

	symbol := m.tape[m.head]
	behavior := m.configurations[m.curr].behaviors[symbol]
	behavior.execute(m)
}

func (m *Machine) resizeTapeLeft() {
	oldLength := len(m.tape)
	oldTape := m.tape
	m.tape = make([]rune, oldLength*2)
	for i := 0; i < oldLength*2; i += 1 {
		m.tape[i] = None
	}
	copy(m.tape[oldLength:], oldTape)
	m.head += oldLength
}

func (m *Machine) resizeTapeRight() {
	oldLength := len(m.tape)
	oldTape := m.tape
	m.tape = make([]rune, oldLength*2)
	for i := 0; i < oldLength*2; i += 1 {
		m.tape[i] = None
	}
	copy(m.tape, oldTape)
}

func (m *Machine) print() {
	for i := len(m.tape) - 1; i >= 0; i -= 1 {
		symbol := m.tape[i]

		if _, ok := m.firstType[symbol]; ok {
			fmt.Printf("%c", symbol)
		}
	}

	fmt.Println()
}
