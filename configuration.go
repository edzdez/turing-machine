package main

type Behavior struct {
	operations         []Operation
	finalConfiguration rune
}

func (b *Behavior) execute(machine *Machine) {
	for _, operation := range b.operations {
		operation.execute(machine)
	}

	machine.curr = b.finalConfiguration
}

type Configuration struct {
	behaviors map[rune]Behavior
}
