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
					Connections: []string{"EH1", "LI2"},
					NumLetters:  4,
				},
				Menu{
					Connections: []string{"LH1", "LI2"},
					NumLetters:  3,
				},
				Menu{
					Connections: []string{"LH1", "OI2"},
					NumLetters:  4,
				},
			},
		},
		{
			input: "WSNPNLKLSTCS",
			crib:  "ATTACKATDAWN",
			result: []Menu{
				Menu{
					Connections: []string{"WA1", "ST2", "NT3", "PA4", "NC5", "LK6", "KA7", "LT8", "SD9", "TA10", "CW11", "SN12"},
					NumLetters:  10,
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
