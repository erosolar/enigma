package menumaker

type Menu struct {
	Connections []string
	NumLetters  int
	graph       map[byte]map[byte]int
}
