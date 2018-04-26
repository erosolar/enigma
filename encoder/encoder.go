package encoder

import "strings"

// Setup returns an enigma object set with the given settings
// including which rotors to use, how they are setup, and the plugboard settings
func Setup(settings Settings) Enigma {
	enig := Enigma{}
	enig.rotors = make([]rotor, 0)
	rotors := settings.RotorOrder
	for idx, rnum := range rotors {
		rstg := settings.RingSettings[idx]
		if rnum < 1 || rnum > 5 ||
			rstg < 1 || rstg > 26 {
			panic("rotor index must be between I and V, and ring setting must be between 1 and 26")
		}

		sub := rotorSubs[rnum-1]

		enig.rotors = append(enig.rotors, rotor{
			substitution: sub[len(sub)-(rstg-1):len(sub)] + sub[0:len(sub)-(rstg-1)],
			turnover:     rotorTurn[rnum-1],
			currPos:      0,
		})

	}
	enig.reflector = settings.Reflector
	return enig
}

//Initialize returns an enigma set up to start encoding using the key 'key'
func Initialize(enig Enigma, key string) Enigma {
	if len(key) != len(enig.rotors) {
		panic("key must be 3 characters long")
	}
	for i := 0; i < len(enig.rotors); i++ {
		enig.rotors[i].currPos = strings.IndexByte(alphabet, key[i])
	}
	return enig
}

func (e *Enigma) Encrypt(s string) string {
	result := ""
	for _, c := range s {
		result += string(e.encryptLetter(c))
	}
	return result
}

func (e *Enigma) encryptLetter(startingLetter rune) byte {
	// TODO plugboard

	// First pass through rotors
	e.stepRotors()

	var letter byte
	currIndex := strings.IndexRune(alphabet, startingLetter)
	for i := 2; i >= 0; i-- {
		currIndex = mod26(currIndex + e.rotors[i].currPos)
		letter = e.rotors[i].substitution[currIndex]
		currIndex = mod26(strings.IndexByte(alphabet, letter) - e.rotors[i].currPos)
	}

	// Reflector
	letter = e.reflector[currIndex]
	currIndex = strings.IndexByte(alphabet, letter)

	// Backwards through rotors
	for i := 0; i < len(e.rotors); i++ {
		currIndex = mod26(currIndex + e.rotors[i].currPos)
		currIndex = mod26(strings.IndexByte(e.rotors[i].substitution, alphabet[currIndex]))
		currIndex = mod26(currIndex - e.rotors[i].currPos)
		letter = alphabet[currIndex]
	}

	return letter
}

func (e *Enigma) stepRotors() {
	if e.rotors[2].currPos == strings.IndexRune(alphabet, e.rotors[2].turnover) {
		e.rotors[1].currPos++
	} else if e.rotors[1].currPos == strings.IndexRune(alphabet, e.rotors[1].turnover) {
		e.rotors[1].currPos++
		e.rotors[0].currPos++
	}
	e.rotors[2].currPos++
}

func mod26(x int) int {
	x = x % 26
	if x < 0 {
		return x + 26
	}
	return x
}
