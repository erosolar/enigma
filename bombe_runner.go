package main

import "github.mit.edu/erosolar/enigma/bombe"
import "github.mit.edu/erosolar/enigma/checker"
import "fmt"

var bombes []bombe.Bombe

var rotorOrders = [][]int{
	[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1},
	[]int{1, 2, 4}, []int{1, 4, 2}, []int{2, 1, 4}, []int{2, 4, 1}, []int{4, 1, 2}, []int{4, 2, 1},
	[]int{1, 2, 5}, []int{1, 5, 2}, []int{2, 1, 5}, []int{2, 5, 1}, []int{5, 1, 2}, []int{5, 2, 1},
	[]int{1, 3, 4}, []int{1, 4, 3}, []int{3, 1, 4}, []int{3, 4, 1}, []int{4, 1, 3}, []int{4, 3, 1},
	[]int{1, 3, 5}, []int{1, 5, 3}, []int{3, 1, 5}, []int{3, 5, 1}, []int{5, 1, 3}, []int{5, 3, 1},
	[]int{1, 4, 5}, []int{1, 5, 4}, []int{4, 1, 5}, []int{4, 5, 1}, []int{5, 1, 4}, []int{5, 4, 1},
	[]int{2, 3, 4}, []int{2, 4, 3}, []int{3, 2, 4}, []int{3, 4, 2}, []int{4, 2, 3}, []int{4, 3, 2},
	[]int{2, 3, 5}, []int{2, 5, 3}, []int{3, 2, 5}, []int{3, 5, 2}, []int{5, 2, 3}, []int{5, 3, 2},
	[]int{2, 4, 5}, []int{2, 5, 4}, []int{4, 2, 5}, []int{4, 5, 2}, []int{5, 2, 4}, []int{5, 4, 2},
	[]int{3, 4, 5}, []int{3, 5, 4}, []int{4, 3, 5}, []int{4, 5, 3}, []int{5, 3, 4}, []int{5, 4, 3},
}

func main() {
	settings := bombe.Settings{
		RotorOrder: []int{3, 2, 1},
		//		Connections: []string{"PA4", "AW1", "WC11", "CN5", "NT3", "TA10", "NS12", "ST2", "SD9", "TL8", "LK6", "KA7"},
		Connections: []string{"CU13", "UD2", "DL5", "LR12", "RA7", "AE1", "ES10", "UK4", "UX16", "XN14", "XN6", "NG15", "GB8"},
		NumEnigmas:  13,
		NumLetters:  13,
	}
	for _, r := range rotorOrders {
		settings.RotorOrder = r
		bombes = append(bombes, bombe.Setup(settings))
	}
	ch := make(chan bombe.Result, 12)
	doneCh := make(chan bool)
	for _, b := range bombes {
		go b.Run('A', 'A', ch, doneCh)
	}

	doneThreads := 0
	results := make([]bombe.Result, 0)
R:
	for {
		select {
		case res := <-ch:
			fmt.Print("*")
			results = append(results, res)
		case <-doneCh:
			doneThreads++
			if doneThreads == 60 { // 60 rotor orders -> 60 bombes in parallel
				fmt.Println()
				fmt.Println("Received all results")
				break R
			}
		}
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
