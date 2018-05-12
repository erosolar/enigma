package menumaker

import "testing"
import "sort"

func Test(t *testing.T) {
	tests := []struct {
		input  string
		crib   string
		result []Menu
	}{
		{
			input: "HELLO",
			crib:  "HI",
			result: []Menu{
				Menu{
					Connections: []string{"EH0", "LI1"},
					NumLetters:  4,
				},
				Menu{
					Connections: []string{"LH0", "LI1"},
					NumLetters:  3,
				},
				Menu{
					Connections: []string{"LH0", "OI1"},
					NumLetters:  4,
				},
			},
		},
	}

	for _, test := range tests {
		m := CreateMenus(test.input, test.crib)
		if !testEqualMenus(m, test.result) {
			t.Error("Expected menu", test.result, "but got", m)
		}
	}
}

func testEqualMenus(m1, m2 []Menu) bool {
	if len(m1) != len(m2) {
		return false
	}

	matched := make(map[int]bool)

	for _, inputMenu := range m1 {
		found := false
		for i, checkMenu := range m2 {
			if !found {
				if ok, _ := matched[i]; ok {
					continue
				}
				if inputMenu.NumLetters != checkMenu.NumLetters {
					continue
				}
				if testEqualConnections(inputMenu.Connections, checkMenu.Connections) {
					found = true
					matched[i] = true
				}
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func testEqualConnections(c1, c2 []string) bool {
	if len(c1) != len(c2) {
		return false
	}

	sort.Strings(c1)
	sort.Strings(c2)

	for i, s := range c1 {
		if s != c2[i] {
			return false
		}
	}

	return true
}
