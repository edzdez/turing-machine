package main

type Operation interface {
	execute(machine *Machine)
}

type P0 struct{}

func (o P0) execute(machine *Machine) {
	machine.tape[machine.head] = Zero
}

type P1 struct{}

func (o P1) execute(machine *Machine) {
	machine.tape[machine.head] = One
}

type PSchwa struct{}

func (o PSchwa) execute(machine *Machine) {
	machine.tape[machine.head] = Schwa
}

type PX struct{}

func (o PX) execute(machine *Machine) {
	machine.tape[machine.head] = X
}

type E struct{}

func (o E) execute(machine *Machine) {
	machine.tape[machine.head] = None
}

type L struct{}

func (o L) execute(machine *Machine) {
	// shift the tape left, which is equivalent to shifting the head right
	machine.head += 1
	if machine.head >= len(machine.tape) {
		machine.resizeTapeRight()
	}
}

type R struct{}

func (o R) execute(machine *Machine) {
	// shift the tape right, which is equivalent to shifting the head left
	machine.head -= 1
	if machine.head <= 0 {
		machine.resizeTapeLeft()
	}
}
