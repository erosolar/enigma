package bombe

type Settings struct {
	RotorOrder []int // top -> bottom left -> right
	NumEnigmas int   // number of edges in the menu (n, below)
	NumLetters int   // number of letters in menu
	// each connection is two capital letters followed by an offset
	// eg. AP12 is an edge between A and P with offset 12
	Connections []string // what's the best way to pass this in? TODO
}

type Bombe struct {
	enig        enigma
	settings    Settings
	state       map[int]map[int]bool // int 1 is char label at node, map is set of things that are 'on'
	connections map[int][]connection // adjacency list of menu connections (indices into edge list)
}

type connection struct {
	transform []int
	endpoints []int
}

type enigma struct {
	rotors [][26]int // substitutions of rotors used
}

// Result defines the result type (a possible setting)
type Result struct {
	Offset    int
	Rotors    []int
	Printable string
	State     map[int]map[int]bool
	Message   string
}
