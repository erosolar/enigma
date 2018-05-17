package zygalski

import "github.mit.edu/erosolar/enigma/encoder"
import "strconv"


const idxA = 65

// returns true if a given setting can return a female
func MakesFemale(fst string, snd string, thd string, en encoder.Enigma) bool {
	key := fst + snd + thd
	for k := 0; k < 26; k++ {
		inp := string(idxA + k) + "ZZ"
		en = encoder.Initialize(en, key)
		ct := en.Encrypt(inp + inp)
		if ct[0] == ct[3] {
			return true
		}
	}
	return false
}

// takes two z-sheets and returns a third which is the result of overlaying them
func ApplyZs(old_zs [][26]bool, new_zs [][26]bool) ([][26]bool, int) {
	count := 0
	out_zs := make([][26]bool, 26)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			out_zs[i][j] = old_zs[i][j] && new_zs[i][j]
			if out_zs[i][j] {
				count++
			}
		}
	}
	return out_zs, count
}

// returns a the first true cell in a z-sheet.  returns -1,-1 if there are not truthy entries
func FindMatch(zs [][26]bool) (int, int) {
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if zs[i][j] {
				return i,j
			}
		}
	}
	return -1, -1
}

// main helper function for trying a particular setting guess for all indicators
func ApplyWheelRing(indicators []string, ring1 int, wo []int) string {
	// set up the z-sheet that will get updated with each additional female
	prev_ind := ""
	main_zs := make([][26]bool,26)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			main_zs[i][j] = true
		}
	}
	cnt := 0
	for _, ind := range indicators {
		ring := []int{ring1, 25, 25}
		if ind[0] != ind[3] {
			continue
		}
		// shift the next sheet based on the previous sheet to line them up
		if prev_ind != "" {
			ring[1] = (25 + int(rune(ind[1])) - int(rune(prev_ind[1]))) % 25
			ring[1] = (25 + int(rune(ind[2])) - int(rune(prev_ind[2]))) % 25
		}
		prev_ind = ind

		settings := encoder.Settings{
			RotorOrder:   wo,
			RingSettings: ring,
			Plugs:        "",
			Reflector:    encoder.UKWB,
		}
		en := encoder.Setup(settings)

		zs := make([][26]bool, 26)

		// form the z-sheet for this indicator
		fst := string(ind[0])
		for snd := 0; snd < 26; snd++ {
			for thd := 0; thd < 26; thd ++ {
				zs[snd][thd] = MakesFemale(fst, string(idxA + snd), string(idxA + thd), en)
			}
		}
		main_zs, cnt = ApplyZs(main_zs, zs)
		if cnt == 0 {
			return ""
		}
	}
	if cnt == 1 {
		// we've found a potential configuration
		mi, mj := FindMatch(main_zs)
		res2 := (25 + int(rune(prev_ind[1])) - mi) % 25
		res3 := (25 + int(rune(prev_ind[2])) - mj) % 25
		return string(ring1 + idxA) + string(res2 + idxA) + string(res3 + idxA)
	}
	return ""
}

// take a list of indicators (the double enciphered key) and find a potential ring settings configuration
func DoZyg(indicators []string) string {
	for _, wo := range []int{123,132,213,231,312,321} {
		for ring1 := 1; ring1 < 26; ring1++ {
			wheels := []int{wo/100, (wo%100)/10, wo%10}
			res := ApplyWheelRing(indicators, ring1, wheels)
			if res != "" {
				return res
			}
		}
	}
	return ""
}
