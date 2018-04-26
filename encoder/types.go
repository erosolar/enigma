package encoder

// for constants
type rotor struct {
	substitution string
	turnover     rune
	ringstellung int // changes position of internal wiring relative to alphabet/turnover
	currPos      int // current letter-index in window
}

type navalRotor struct {
	substitution string
	turnover     [2]rune
	ringstellung int // changes position of internal wiring relative to alphabet/turnover
}

type Settings struct {
	RotorOrder   []int
	RingSettings []int
	// Plugs        [][2]rune
	Reflector    string
}

type Enigma struct {
	rotors    []rotor
	plugboard string
	reflector string
}
