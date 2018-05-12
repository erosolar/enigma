package menumaker

type Menu struct {
	Connections []string
	NumLetters  int
	letters     map[byte]bool
}
