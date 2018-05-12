package checker

func CheckIfPossiblePlugboard(input map[int]map[int]bool) bool {
	oneLitUp := false
	allButOneLitUp := false
	for _, m := range input {
		if len(m) == 1 {
			oneLitUp = true
		} else if len(m) == 25 {
			allButOneLitUp = true
		} else {
			return false
		}
	}

	if oneLitUp && allButOneLitUp {
		return false
	}

	stekers := make(map[int]bool)
	if oneLitUp {
		for _, m := range input {
			for s, _ := range m {
				if _, ok := stekers[s]; ok {
					return false
				} else {
					stekers[s] = true
				}
			}
		}
	} else {
		for _, m := range input {
			all := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true, 20: true, 21: true, 22: true, 23: true, 24: true, 25: true}
			for s, _ := range m {
				delete(all, s)
			}
			for s, _ := range all {
				if _, ok := stekers[s]; ok {
					return false
				} else {
					stekers[s] = true
				}
			}
		}
	}

	return true
}