package main

const (
	Zero  rune = '0'
	One        = '1'
	None       = ' '
	Schwa      = 'ə'
	X          = 'x'
)

var p0 *P0 = &P0{}
var p1 *P1 = &P1{}
var pSchwa *PSchwa = &PSchwa{}
var pX *PX = &PX{}
var e *E = &E{}
var l *L = &L{}
var r *R = &R{}

func example1() *Machine {
	configurations := map[rune]Configuration{
		'b': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{p0, r},
					finalConfiguration: 'c',
				},
			},
		},
		'c': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{r},
					finalConfiguration: 'e',
				},
			},
		},
		'e': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{p1, r},
					finalConfiguration: 'f',
				},
			},
		},
		'f': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{r},
					finalConfiguration: 'b',
				},
			},
		},
	}

	symbolsOfTheFirstType := map[rune]bool{
		Zero: true,
		One:  true,
	}

	return NewMachine(configurations, 'b', symbolsOfTheFirstType)
}

func example1Condensed() *Machine {
	configurations := map[rune]Configuration{
		'b': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{p0},
					finalConfiguration: 'b',
				},
				Zero: {
					operations:         []Operation{r, r, p1},
					finalConfiguration: 'b',
				},
				One: {
					operations:         []Operation{r, r, p0},
					finalConfiguration: 'b',
				},
			},
		},
	}

	symbolsOfTheFirstType := map[rune]bool{
		Zero: true,
		One:  true,
	}

	return NewMachine(configurations, 'b', symbolsOfTheFirstType)
}

func example2() *Machine {
	configurations := map[rune]Configuration{
		'b': {
			behaviors: map[rune]Behavior{
				None: {
					operations:         []Operation{pSchwa, r, pSchwa, r, p0, r, r, p0, l, l},
					finalConfiguration: 'o',
				},
			},
		},
		'o': {
			behaviors: map[rune]Behavior{
				One: {
					operations:         []Operation{r, pX, l, l, l},
					finalConfiguration: 'o',
				},
				Zero: {
					operations:         []Operation{},
					finalConfiguration: 'q',
				},
			},
		},
		'q': {
			behaviors: map[rune]Behavior{
				One: {
					operations:         []Operation{r, r},
					finalConfiguration: 'q',
				},
				Zero: {
					operations:         []Operation{r, r},
					finalConfiguration: 'q',
				},
				None: {
					operations:         []Operation{p1, l},
					finalConfiguration: 'p',
				},
			},
		},
		'p': {
			behaviors: map[rune]Behavior{
				X: {
					operations:         []Operation{e, r},
					finalConfiguration: 'q',
				},
				Schwa: {
					operations:         []Operation{r},
					finalConfiguration: 'f',
				},
				None: {
					operations:         []Operation{l, l},
					finalConfiguration: 'p',
				},
			},
		},
		'f': {
			behaviors: map[rune]Behavior{
				Zero: {
					operations:         []Operation{r, r},
					finalConfiguration: 'f',
				},
				One: {
					operations:         []Operation{r, r},
					finalConfiguration: 'f',
				},
				X: {
					operations:         []Operation{r, r},
					finalConfiguration: 'f',
				},
				Schwa: {
					operations:         []Operation{r, r},
					finalConfiguration: 'f',
				},
				None: {
					operations:         []Operation{p0, l, l},
					finalConfiguration: 'o',
				},
			},
		},
	}

	symbolsOfTheFirstType := map[rune]bool{
		Zero: true,
		One:  true,
	}

	return NewMachine(configurations, 'b', symbolsOfTheFirstType)
}

func main() {
	machine := example1()
	for i := 0; i < 200; i += 1 {
		machine.step()
	}
	machine.print()

	machine = example1Condensed()
	for i := 0; i < 100; i += 1 {
		machine.step()
	}
	machine.print()

	machine = example2()
	for i := 0; i < 2475; i += 1 {
		machine.step()
	}
	machine.print()
}
