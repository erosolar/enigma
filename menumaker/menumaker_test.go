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
			input: "WSNPNLKLSTCS",
			crib:  "ATTACKATDAWN",
			result: []Menu{
				Menu{
					Connections: []string{"AW1", "ST2", "NT3", "AP4", "CN5", "KL6", "AK7", "LT8", "DS9", "AT10", "CW11", "NS12"},
					NumLetters:  10,
				},
			},
		},
		{
			input: "QFZWRWIVTYRESXBFOGKUHQBAISEZ",
			crib:  "WETTERVORHERSAGEBISKAYA",
			result: []Menu{
				Menu{
					Connections: []string{"HX14", "BH21", "BS23", "BE15", "SY26", "RS13", "OS17", "RY10", "FR16", "RV11", "RW5", "EW6", "EO12", "AE27", "ET9", "EU20", "AG18", "AK24", "GK19", "AI25", "IQ22", "IT7", "TV8"},
					NumLetters:  18,
				},
			},
		},
	}

	for _, test := range tests {
		m := MakeMenus(test.input, test.crib)
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
