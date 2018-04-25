package encoder

// for constants
type rotor struct {
	substitution string
	turnover     rune
	ringstellung int // changes position of internal wiring relative to alphabet/turnover
}

type navalRotor struct {
	substitution string
	turnover     [2]rune
	ringstellung int // changes position of internal wiring relative to alphabet/turnover
}

type Settings struct {
	rotorOrder   []int
	ringSettings []int
	plugs        [][2]rune
}
