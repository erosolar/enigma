package encoder

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// moving rotors
const (
	// original three rotors
	rotorIsubs   = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotorIturn   = 'Q'
	rotorIIsubs  = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	rotorIIturn  = 'E'
	rotorIIIsubs = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	rotorIIIturn = 'V'
	// added 1938 (basically original)
	rotorIVsubs = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
	rotorIVturn = 'J'
	rotorVsubs  = "VZBRGITYUPSDNHLXAWMJQOFECK"
	rotorVturn  = 'Z'
	// naval rotors
	rotorVIsubs    = "JPGVOUMFYQBENHZRDKASXLICTW"
	rotorVIIsubs   = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
	rotorVIIIsubs  = "EJMZALYXVBWFCRQUONTSPIKHGD"
	navalRotorTurn = "ZM"
)

var rotorSubs = [5]string{
	rotorIsubs,
	rotorIIsubs,
	rotorIIIsubs,
	rotorIVsubs,
	rotorVsubs,
}
var rotorTurn = [5]rune{
	rotorIturn,
	rotorIIturn,
	rotorIIIturn,
	rotorIVturn,
	rotorVturn,
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
