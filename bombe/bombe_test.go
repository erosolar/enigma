package bombe

import "testing"
import "fmt"

func TestSetup(t *testing.T) {
	fmt.Println("setting up bombe...")
	settings := Settings{
		RotorOrder:  []int{1, 2, 3},
		NumEnigmas:  2,
		NumLetters:  3,
		Connections: []string{"AB2", "AC12"},
	}

	bombe := Setup(settings)
	t.Log("bombe should have correct enigma settings")
	if bombe.enig.rotors[0] != constRotors[0] ||
		bombe.enig.rotors[1] != constRotors[1] ||
		bombe.enig.rotors[2] != constRotors[2] {
		t.FailNow()
	}
	settings.RotorOrder = []int{4, 5, 2}
	t.Log("should work with a different ordering")
	bombe = Setup(settings)
	if bombe.enig.rotors[0] != constRotors[3] ||
		bombe.enig.rotors[1] != constRotors[4] ||
		bombe.enig.rotors[2] != constRotors[1] {
		t.FailNow()
	}
}

func TestMakeSystem(t *testing.T) {
	fmt.Println("making a system...")
	settings := Settings{
		RotorOrder:  []int{1, 2, 3},
		NumEnigmas:  2,
		NumLetters:  3,
		Connections: []string{"AB2", "AC12"},
	}

	bombe := Setup(settings)

	bombe.makeSystem(0)
	t.Log("after making system, state and connections should not be nil")
	if bombe.state == nil ||
		bombe.connections == nil {
		t.FailNow()
	}
	t.Log("after making system, state and connections should have entries for each letter passed in")
	if _, ok := bombe.state[0]; !ok { // 'A' = 0
		t.FailNow()
	}
	if _, ok := bombe.state[1]; !ok { // 'B' = 1
		t.FailNow()
	}
	if _, ok := bombe.state[2]; !ok { // 'C' = 2
		t.FailNow()
	}
	if _, ok := bombe.connections[0]; !ok {
		t.FailNow()
	}
	if _, ok := bombe.connections[1]; !ok {
		t.FailNow()
	}
	if _, ok := bombe.connections[2]; !ok {
		t.FailNow()
	}
	t.Log("after making system, states should be empty")
	if len(bombe.state[0]) != 0 ||
		len(bombe.state[1]) != 0 ||
		len(bombe.state[2]) != 0 {
		t.FailNow()
	}
	t.Log("after making system, connections should be correct")
	if len(bombe.connections[0]) != 2 ||
		len(bombe.connections[1]) != 1 ||
		len(bombe.connections[2]) != 1 {
		t.FailNow()
	}
	t.Log("connections should make sense")
	c := bombe.connections[1][0]
	if (c.endpoints[0] != 0 && c.endpoints[0] != 1) ||
		(c.endpoints[1] != 0 && c.endpoints[1] != 1) {
		t.FailNow()
	}
	t.Log("connection transform should be correct")
	t.Log(c.transform)
	t.Log(bombe.enig.makeTransform(2))
	for i, r := range bombe.enig.makeTransform(2) {
		if c.transform[i] != r {
			t.FailNow()
		}
	}
}

func TestEncryptLetter(t *testing.T) {
	settings := Settings{
		RotorOrder:  []int{3, 2, 1},
		NumEnigmas:  2,
		NumLetters:  3,
		Connections: []string{"AB2", "AC12"},
	}

	bombe := Setup(settings)

	t.Log("A -> B")
	res := bombe.enig.encryptLetter(0, 1) // 'A'
	if res != 1 {
		t.Errorf("A -> %v", rune(res))
	}
	t.Log("A -> D")
	res = bombe.enig.encryptLetter(0, 2) // 'A' again
	if res != 3 {
		t.Errorf("A -> %v", rune(res))
	}
	t.Log("A -> Z")
	res = bombe.enig.encryptLetter(0, 3) // 'A' again
	if res != 25 {
		t.Errorf("A -> %v", rune(res))
	}
	t.Log("A -> G")
	res = bombe.enig.encryptLetter(0, 4) // 'A' again
	if res != 6 {
		t.Errorf("A -> %v", rune(res))
	}
	t.Log("A -> O")
	res = bombe.enig.encryptLetter(0, 5) // 'A' again
	if res != 14 {
		t.Errorf("A -> %v", rune(res))
	}

}

func TestFindSteadyState(t *testing.T) {
	settings := Settings{
		RotorOrder:  []int{3, 2, 1},
		Connections: []string{"PA4", "AW1", "WC11", "CN5", "NT3", "TA10", "NS12", "ST2", "SD9", "TL8", "LK6", "KA7"},
		NumEnigmas:  12,
		NumLetters:  10,
	}

	bombe := Setup(settings)
	bombe.makeSystem(0)
	bombe.initialize()
}
