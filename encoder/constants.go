package encoder

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// original three rotors
var rotorI = rotor{
	substitution: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	turnover:     'Q', // Q->R = turnover of next rotor
}
var rotorII = rotor{
	substitution: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	turnover:     'E',
}
var rotorIII = rotor{
	substitution: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	turnover:     'V',
}

// added 1938 (basically original)
var rotorIV = rotor{
	substitution: "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	turnover:     'J',
}
var rotorV = rotor{
	substitution: "VZBRGITYUPSDNHLXAWMJQOFECK",
	turnover:     'Z',
}

// naval rotors
var rotorVI = navalRotor{
	substitution: "JPGVOUMFYQBENHZRDKASXLICTW",
	turnover:     [2]rune{'Z', 'M'},
}
var rotorVII = navalRotor{
	substitution: "NZJHGRCXMYSWBOUFAIVLPEKQDT",
	turnover:     [2]rune{'Z', 'M'},
}
var rotorVIII = navalRotor{
	substitution: "FKQHTLXOCBJSPDZRAMEWNIUYGV",
	turnover:     [2]rune{'Z', 'M'},
}

// still rotors
const (
	// entry wheel
	ETW = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// used until 1 nov 1937
	UKWA = "EJMZALYXVBWFCRQUONTSPIKHGD"
	// used starting 1 nov 1937
	UKWB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	// used for a short period in 1940 (possibly by mistake?)
	// uncle walter!
	UKWC = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
)
