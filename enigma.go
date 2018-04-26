package main

import (
	"fmt"

	"github.mit.edu/erosolar/enigma/encoder"
)

func main() {
	fmt.Printf("Welcome to Enigma.\n")

	settings := encoder.Settings{
		RotorOrder:   []int{1, 2, 3},
		RingSettings: []int{1, 1, 1},
		Plugs:        "",
		Reflector:    encoder.UKWB,
	}

	enigma := encoder.Setup(settings)
	enigma = encoder.Initialize(enigma, "AAA")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be BDZGO)")

	enigma = encoder.Initialize(enigma, "ADU")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be EQIBM)")

	enigma = encoder.Initialize(enigma, "ZDU")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be CPQZG)")

	fmt.Println("WITH RING SETTINGS BBB")
	settings.RingSettings = []int{2, 2, 2}
	enigma = encoder.Setup(settings)
	enigma = encoder.Initialize(enigma, "AAA")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be EWTYX)")
	enigma = encoder.Initialize(enigma, "ADU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be ECQEZ)")
	enigma = encoder.Initialize(enigma, "ZDU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be QQPSG)")
	fmt.Println("WITH RING SETTINGS BBB and plugboard 'PO ML IU KJ NH YT GB VF RE DC'")
	settings.RingSettings = []int{2, 2, 2}
	settings.Plugs = "PO ML IU KJ NH YT GB VF RE DC"
	enigma = encoder.Setup(settings)
	enigma = encoder.Initialize(enigma, "AAA")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be RWYTX)")
	enigma = encoder.Initialize(enigma, "ADU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be RDQRZ)")
	enigma = encoder.Initialize(enigma, "ZDU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be QQOSB)")
}
