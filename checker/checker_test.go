package checker

import (
	"sort"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input     map[int]map[int]bool
		plugboard string
		legal     bool
	}{
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{19: true}, 15: map[int]bool{6: true}, 10: map[int]bool{22: true}, 11: map[int]bool{9: true}, 19: map[int]bool{8: true}, 13: map[int]bool{18: true}, 18: map[int]bool{7: true}, 3: map[int]bool{2: true}, 2: map[int]bool{0: true}, 22: map[int]bool{21: true},
		// 	},
		// 	plugboard: "",
		// 	legal:     false,
		// },
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{6: true}, 15: map[int]bool{25: true}, 10: map[int]bool{23: true}, 11: map[int]bool{7: true}, 19: map[int]bool{9: true}, 13: map[int]bool{5: true}, 18: map[int]bool{24: true}, 3: map[int]bool{21: true}, 2: map[int]bool{16: true}, 22: map[int]bool{8: true},
		// 	},
		// 	plugboard: "AG PZ KX HL JT FN SY DV CQ IW",
		// 	legal:     true,
		// },
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{7: true}, 15: map[int]bool{21: true}, 10: map[int]bool{10: true}, 11: map[int]bool{20: true}, 19: map[int]bool{25: true}, 13: map[int]bool{12: true}, 18: map[int]bool{4: true}, 3: map[int]bool{24: true}, 2: map[int]bool{9: true}, 22: map[int]bool{1: true},
		// 	},
		// 	plugboard: "AH PV KK LU TZ MN ES DY CJ BW",
		// 	legal:     true,
		// },
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{0: true}, 15: map[int]bool{15: true}, 10: map[int]bool{10: true}, 11: map[int]bool{11: true}, 19: map[int]bool{19: true}, 13: map[int]bool{13: true}, 18: map[int]bool{18: true}, 3: map[int]bool{23: true}, 2: map[int]bool{25: true}, 22: map[int]bool{21: true},
		// 	},
		// 	plugboard: "AA PP KK LL TT NN SS DX CZ VW",
		// 	legal:     false,
		// },
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{0: true, 1: true, 2: true, 3: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 15: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true}, 10: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 11: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 19: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 24: true, 25: true}, 13: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 25: true}, 18: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 3: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 2: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 22: map[int]bool{0: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true},
		// 	},
		// 	plugboard: "AE PZ KR LQ TX NY JS DO CG BW",
		// 	legal:     true,
		// },
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 15: map[int]bool{0: true, 1: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 10: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 11: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 22: true, 23: true, 24: true, 25: true}, 19: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true}, 13: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 18: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 3: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 25: true}, 2: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 22: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 23: true, 24: true, 25: true},
		// 	},
		// 	plugboard: "AK CP KM LV TZ LN RS DY CJ WW",
		// 	legal:     false,
		// },
		{
			input: map[int]map[int]bool{
				0: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 15: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true}, 10: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 11: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 19: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 13: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 18: map[int]bool{0: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 3: map[int]bool{0: true, 1: true, 2: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 2: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 22: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 24: true, 25: true},
			},
			plugboard: "AR PZ KK LL GT NU BS DD CQ WX",
			legal:     true,
		},
		// {
		// 	input: map[int]map[int]bool{
		// 		0: map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 15: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 10: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 11: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 19: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 13: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 18: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 3: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 2: map[int]bool{0: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}, 22: map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true},
		// 	},
		// 	plugboard: "AA PP KK LL TT NN SS DU BC HW",
		// 	legal:     false,
		// },
	}

	for i, test := range tests {
		pb, l := CheckIfPossiblePlugboard(test.input)
		if l != test.legal {
			t.Error(i, "Expected legal", test.legal, "but got", l)
		} else if test.legal && !testEqualPlugboards(pb, test.plugboard) {
			t.Error(i, "Expected plugboard", test.plugboard, "but got", pb)
		}
	}
}

func testEqualPlugboards(p1, p2 string) bool {
	c1 := strings.Fields(p1)
	c2 := strings.Fields(p2)

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
