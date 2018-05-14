package bombe

import "strconv"
import "fmt"
import "time"

//GetResults takes in a settings object
// if RotorOrder is set, does that rotor order
// else, tries all possible rotor orders
// nthreads is how many threads to run in parallel
// resChan will be closed once there are no more results
// IMPORTANT: this function is intended to be run as a goroutine!!!
func GetResults(settings Settings, nThreads int, resChan chan Result, killCh chan bool) {
	var bombes []Bombe
	if settings.RotorOrder != nil {
		bombes = make([]Bombe, 0, 1)
		bombes[0] = Setup(settings)
	} else {
		bombes = make([]Bombe, 0, len(rotorOrders))
		for _, r := range rotorOrders {
			settings.RotorOrder = r
			bombes = append(bombes, Setup(settings))
		}
	}
	fmt.Printf("Running %d parallel threads on %d bombes\n", nThreads, len(bombes))
	start := time.Now()
	jobs := make(chan Bombe, nThreads)
	ch := make(chan Result, 12)
	doneCh := make(chan bool)
	for i := 0; i < nThreads; i++ {
		go threadRun(jobs, ch, doneCh)
	}

	doneThreads := 0
	go func() {
		for _, b := range bombes {
			jobs <- b
		}
	}()
R:
	for {
		select {
		case res := <-ch:
			resChan <- res
		case <-doneCh:
			doneThreads++
			if doneThreads == nThreads {
				fmt.Println()
				fmt.Println("Received all results")
				fmt.Printf("bombes took: %s\n", time.Since(start))
				break R
			}
		case _, ok := <-killCh:
			if !ok {
				break R
			}
		}
	}
	close(resChan)
}

func threadRun(jobCh chan Bombe, resultCh chan Result, doneCh chan bool) {
	done := make(chan bool)
	for {
		select {
		case b := <-jobCh:
			b.Run(resultCh, done)
		case <-time.After(time.Second):
			doneCh <- true
			return
		}
	}
}

// Setup initializes the Bombe to have the provided rotors and connections
func Setup(settings Settings) Bombe {
	b := Bombe{}
	b.settings = settings
	// make enigma clone
	rotors := make([][26]int, 0, 3)
	for _, j := range settings.RotorOrder {
		rotors = append(rotors, constRotors[j-1])
	}
	b.enig = enigma{
		rotors: rotors,
	}
	return b
}

// Run runs the bombe with a given starting guess
// pt: the node you want to give a starting guess for
// plugboard: the guess for the node
func (b Bombe) Run(resultChan chan Result, doneChan chan bool) {
	for offset := 0; offset < 26*26*26; offset++ {
		b.makeSystem(offset)
		pt := b.initialize()
		b.findSteadyState(pt)
		if offset%(26*26) == 0 {
			fmt.Printf(".")
		}
		if testOutput(b.state) {
			resultChan <- Result{
				Offset:    offset,
				Rotors:    b.settings.RotorOrder,
				Printable: formatOutput(b.state),
				State:     b.state,
			}
		}
	}
}

// assumes initial offset of 0
func (b *Bombe) makeSystem(offset int) {
	connections := make(map[int][]connection, b.settings.NumLetters)
	state := make(map[int]map[int]bool, b.settings.NumLetters)
	for _, conn := range b.settings.Connections {
		n1 := int(conn[0] - 'A')
		n2 := int(conn[1] - 'A')
		off, err := strconv.Atoi(conn[2:])
		if err != nil {
			panic("offset must be number; given " + conn[2:])
		}
		// make sure letters exist in state things
		if _, ok := state[n1]; !ok {
			state[n1] = make(map[int]bool)
		}
		if _, ok := state[n2]; !ok {
			state[n2] = make(map[int]bool)
		}
		if _, ok := connections[n1]; !ok {
			connections[n1] = make([]connection, 0)
		}
		if _, ok := connections[n2]; !ok {
			connections[n2] = make([]connection, 0)
		}

		c := connection{
			transform: b.enig.makeTransform(off + offset),
			endpoints: []int{n1, n2},
		}
		connections[n1] = append(connections[n1], c)
		connections[n2] = append(connections[n2], c)
	}
	b.connections = connections
	b.state = state
}

func (b *Bombe) initialize() int {
	// v janky :/
	for k := range b.state {
		b.state[k][k] = true
		return k
	}
	panic("nothing in state!!!")
}

func (b *Bombe) findSteadyState(start int) {
	queue := []int{start}
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		edges := b.connections[elem]
		for _, edge := range edges {
			for _, endp := range edge.endpoints {
				if endp == elem {
					continue
				}
				newSt := transform(b.state[elem], edge.transform)
				if !same(b.state[endp], newSt) {
					for k := range newSt {
						b.state[endp][k] = true
					}
					if !exists(endp, queue) {
						queue = append(queue, endp)
					}
				}
			}
		}
		var newNodes []int
		newNodes, b.state = diagonalBoard(b.state)
		for _, node := range newNodes {
			if !exists(node, queue) {
				queue = append(queue, node)
			}
		}
	}
}

func formatOutput(input map[int]map[int]bool) string {
	out := ""
	for pt, m := range input {
		letters := make([]rune, 26)
		for i := range letters {
			letters[i] = rune(i + 'A')
		}
		for ct := range m {
			letters[ct] = '_'
		}
		out += string(letters)
		out += fmt.Sprintf(" <- %q\n", rune(pt+'A'))
	}
	return out
}

func testOutput(input map[int]map[int]bool) bool {
	count := 0
	for _, m := range input {
		count += len(m)
	}
	if count < 26*len(input) {
		return true
	}
	return false
}

func transform(input map[int]bool, transform []int) map[int]bool {
	output := make(map[int]bool, len(input))
	for k := range input {
		output[transform[k]] = true
	}
	return output
}

func exists(elem int, list []int) bool {
	for _, x := range list {
		if x == elem {
			return true
		}
	}
	return false
}

func same(map1, map2 map[int]bool) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	for k, v := range map2 {
		if map1[k] != v {
			return false
		}
	}
	return true
}

func diagonalBoard(state map[int]map[int]bool) ([]int, map[int]map[int]bool) {
	out := make(map[int]map[int]bool, len(state))
	toCheck := make([]int, 0)
	for k, v := range state {
		out[k] = make(map[int]bool, len(v))
	}
	for k, v := range state {
		for k2 := range v {
			out[k][k2] = true
			if _, ok := state[k2]; ok {
				if _, ok := state[k2][k]; !ok {
					out[k2][k] = true
					toCheck = append(toCheck, k2)
				}
			}
		}
	}
	return toCheck, out
}

// letters go in the top rotor, through each rotor, through the ukw, then back
// through all three rotors in reverse order
func (e enigma) makeTransform(offset int) []int {
	transform := make([]int, 26)
	for i := range transform {
		transform[i] = e.encryptLetter(i, offset)
	}
	return transform
}

func (e enigma) encryptLetter(letter, off int) int {
	offsets := []int{off % 26, (off % (26 * 26)) / 26, off / (26 * 26)}
	idx := letter
	// first pass through
	for i := 0; i < 3; i++ {
		idx = mod26(idx + offsets[i])
		idx = e.rotors[i][idx]
		idx = mod26(idx - offsets[i])
	}
	//reflector
	idx = ukwb[idx]
	// backwards
	for i := 2; i >= 0; i-- {
		idx = mod26(idx + offsets[i])
		for index, elem := range e.rotors[i] {
			if elem == idx {
				idx = index
				break
			}
		}
		idx = mod26(idx - offsets[i])
	}
	return idx
}

func mod26(x int) int {
	x = x % 26
	if x < 0 {
		return x + 26
	}
	return x
}
