package encoder

// Setup returns an enigma object set with the given settings
// including which rotors to use, how they are setup, and the plugboard settings
func Setup(settings Settings) Enigma {
	enig := Enigma{}
	enig.rotors = make([]rotor, 3)
	rotors := settings.rotorOrder
	for idx, rnum := range rotors {
		rstg := settings.ringSettings[idx]
		if rnum < 1 || rnum > 5 ||
			rstg < 1 || rstg > 26 {
			panic("rotor index must be between I and V, and ring setting must be between 1 and 26")
		}

		enig.rotors = append(enig.rotors, rotor{
			substitution: rotorSubs[rnum-1],
			turnover:     rotorTurn[rnum-1],
			ringstellung: rstg - 1,
			currPos:      0,
		})

	}
	return enig
}

//Initialize returns an enigma set up to start encoding using the key 'key'
func Initialize(enig Enigma, key string) Enigma {
	if len(key) != 3 {
		panic("key must be 3 characters long")
	}
	return enig
}
