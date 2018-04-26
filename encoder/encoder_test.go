package encoder

import (
	"fmt"
	"testing"
)

func TestScrambler(t *testing.T) {
	fmt.Println("regular settings... ")
	settings := Settings{
		RotorOrder:   []int{1, 2, 3},
		RingSettings: []int{1, 1, 1},
		Plugs:        "",
		Reflector:    UKWB,
	}
	enigma := Setup(settings)
	enigma = Initialize(enigma, "AAA")
	t.Log("encrypting 'AAAAA' with key 'AAA'")
	if enigma.Encrypt("AAAAA") != "BDZGO" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ADU")
	t.Log("encrypting 'AAAAA' with key 'ADU'")
	if enigma.Encrypt("AAAAA") != "EQIBM" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ZDU")
	t.Log("encrypting 'AAAAA' with key 'ZDU'")
	if enigma.Encrypt("AAAAA") != "CPQZG" {
		t.Fail()
	}
}

func TestRingSettings(t *testing.T) {
	fmt.Println("with ring settings... ")
	settings := Settings{
		RotorOrder:   []int{1, 2, 3},
		RingSettings: []int{2, 2, 2},
		Plugs:        "",
		Reflector:    UKWB,
	}
	enigma := Setup(settings)
	enigma = Initialize(enigma, "AAA")
	t.Log("encrypting 'AAAAA' with key 'AAA'")
	if enigma.Encrypt("AAAAA") != "EWTYX" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ADU")
	t.Log("encrypting 'AAAAA' with key 'ADU'")
	if enigma.Encrypt("AAAAA") != "ECQEZ" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ZDU")
	t.Log("encrypting 'AAAAA' with key 'ZDU'")
	if enigma.Encrypt("AAAAA") != "QQPSG" {
		t.Fail()
	}
}

func TestPlugboard(t *testing.T) {
	fmt.Println("with plugboard... ")
	settings := Settings{
		RotorOrder:   []int{1, 2, 3},
		RingSettings: []int{2, 2, 2},
		Plugs:        "PO ML IU KJ NH YT GB VF RE DC",
		Reflector:    UKWB,
	}
	enigma := Setup(settings)
	enigma = Initialize(enigma, "AAA")
	t.Log("encrypting 'AAAAA' with key 'AAA'")
	if enigma.Encrypt("AAAAA") != "RWYTX" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ADU")
	t.Log("encrypting 'AAAAA' with key 'ADU'")
	if enigma.Encrypt("AAAAA") != "RDQRZ" {
		t.Fail()
	}
	enigma = Initialize(enigma, "ZDU")
	t.Log("encrypting 'AAAAA' with key 'ZDU'")
	if enigma.Encrypt("AAAAA") != "QQOSB" {
		t.Fail()
	}
}

/*
func main() {
	settings.Plugs = "PO ML IU KJ NH YT GB VF RE DC"
	enigma = encoder.Setup(settings)
	enigma = encoder.Initialize(enigma, "AAA")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be RWYTX)")
	enigma = encoder.Initialize(enigma, "ADU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be RDQRZ)")
	enigma = encoder.Initialize(enigma, "ZDU")
	fmt.Println("Encryption of 'AAAAA' is", enigma.Encrypt("AAAAA"), "(should be QQOSB)")
}
*/
