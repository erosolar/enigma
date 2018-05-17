package menumaker

import "strconv"

func MakeMenus(input, crib string) []Menu {
	var menus []Menu
	// check every possible place the crib could be
	for i := 0; i <= len(input)-len(crib); i++ {
		if ok, m := makeMenu(i, input[i:i+len(crib)], crib); ok {
			menus = append(menus, m)
		}
	}

	return menus
}

func makeMenu(start int, input, crib string) (bool, Menu) {
	m := Menu{
		Connections: []string{},
		NumLetters:  0,
		graph:       make(map[byte]map[byte]int),
	}
	for j, s := range input {
		if byte(s) == crib[j] {
			return false, m
		}
		m.addConnection(s, crib[j], start+j)
	}

	m.reformat(crib)

	return true, m
}

func (m *Menu) addConnection(a rune, b byte, pos int) {
	if _, ok := m.graph[byte(a)]; ok {
		m.graph[byte(a)][b] = pos
	} else {
		m.graph[byte(a)] = map[byte]int{b: pos}
	}
	if _, ok := m.graph[b]; ok {
		m.graph[b][byte(a)] = pos
	} else {
		m.graph[b] = map[byte]int{byte(a): pos}
	}
}

func (m *Menu) reformat(crib string) {
	var mainGraph []byte
	allAddedLetters := make(map[byte]bool)

	for _, c := range crib {
		if _, ok := allAddedLetters[byte(c)]; ok {
			continue
		}
		nodes := []byte{}
		queue := []byte{byte(c)}
		allAddedLetters[byte(c)] = true
		var letter byte
		run := true
		for run {
			letter, queue = queue[0], queue[1:]
			nodes = append(nodes, letter)
			for newLetter, _ := range m.graph[letter] {
				if _, ok := allAddedLetters[newLetter]; !ok {
					queue = append(queue, newLetter)
					allAddedLetters[newLetter] = true
				}
			}
			run = len(queue) > 0
		}

		if len(nodes) > len(mainGraph) {
			mainGraph = nodes
		}
	}

	m.NumLetters = 0
	for _, a := range mainGraph {
		m.NumLetters++
		for b, pos := range m.graph[a] {
			if a < b {
				connectionString := string(a) + string(b) + strconv.Itoa(pos+1)
				m.Connections = append(m.Connections, connectionString)
			}
		}
	}
}
