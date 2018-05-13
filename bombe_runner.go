package main

import "github.mit.edu/erosolar/enigma/bombe"
import "github.mit.edu/erosolar/enigma/checker"
import "fmt"

var bombes []bombe.Bombe

func main() {
	settings := bombe.Settings{
		//		Connections: []string{"PA4", "AW1", "WC11", "CN5", "NT3", "TA10", "NS12", "ST2", "SD9", "TL8", "LK6", "KA7"},
		Connections: []string{"CU13", "UD2", "DL5", "LR12", "RA7", "AE1", "ES10", "UK4", "UX16", "XN14", "XN6", "NG15", "GB8"},
		NumEnigmas:  13,
		NumLetters:  13,
	}
	resChan := make(chan bombe.Result)
	go bombe.GetResults(settings, 3, resChan)

	results := make([]bombe.Result, 0)

R:
	for {
		res, ok := <-resChan
		if !ok {
			break R
		}
		results = append(results, res)
	}

	// now for the human checking
	for _, res := range results {
		if checker.CheckIfPossiblePlugboard(res.State) {
			fmt.Println()
			fmt.Printf("offset: %v; rotors: %v\n", res.Offset, res.Rotors)
			fmt.Println(res.Printable)
		}
	}
}
