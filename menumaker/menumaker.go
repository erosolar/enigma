package menumaker

import "strconv"

func MakeMenus(input, crib string) []Menu {
	var menus []Menu
	// check every possible place the crib could be
	for i := 0; i <= len(input)-len(crib); i++ {
		if ok, m := makeMenu(input[i:i+len(crib)], crib); ok {
			menus = append(menus, m)
		}
	}

	return menus
}

func makeMenu(input, crib string) (bool, Menu) {
	m := Menu{
		Connections: []string{},
		NumLetters:  0,
		letters:     make(map[byte]bool),
	}
	for j, s := range input {
		if byte(s) == crib[j] {
			return false, m
		}
		m.addConnection(s, crib[j], j)
	}
	return true, m
}

func (m *Menu) addConnection(a rune, b byte, pos int) {
	connectionString := string(a) + string(b) + strconv.Itoa(pos+1)
	m.Connections = append(m.Connections, connectionString)

	for _, l := range []byte{byte(a), b} {
		if _, ok := m.letters[l]; !ok {
			m.letters[l] = true
			m.NumLetters++
		}
	}
}
