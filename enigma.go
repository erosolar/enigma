package main

import (
	"fmt"

	"github.mit.edu/erosolar/enigma/encoder"
)

func main() {
	fmt.Printf("Welcome to Enigma.\n")

	settings := encoder.Settings{
		[]int{1, 2, 3},
		[]int{1, 1, 1},
		encoder.UKWB,
	}

	enigma := encoder.Setup(settings)
	enigma = encoder.Initialize(enigma, "AAA")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be BDZGO)")

	enigma = encoder.Initialize(enigma, "ADU")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be EQIBM)")

	enigma = encoder.Initialize(enigma, "ZDU")

	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be CPQZG)")

    settings.RingSettings = []int{2,2,2}
    enigma = encoder.Setup(settings)
    enigma = encoder.Initialize(enigma, "AAA")
    fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be EWTYX)")
}
