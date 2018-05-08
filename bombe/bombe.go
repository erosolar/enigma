package bombe

import "strconv"
import "fmt"

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
func (b Bombe) Run(pt, plugboard rune, resultChan chan Result, doneChan chan bool) {
	for offset := 0; offset < 26*26*26; offset++ {
		b.makeSystem(offset)
		b.initialize(pt, plugboard)
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
	doneChan <- true
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

func (b *Bombe) initialize(pt, ct rune) {
	ptInt := int(pt - 'A')
	ctInt := int(ct - 'A')
	b.state[ptInt][ctInt] = true
}

func (b *Bombe) findSteadyState(start rune) {
	queue := []int{int(start - 'A')}
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
