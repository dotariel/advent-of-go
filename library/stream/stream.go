package stream

func Parse(input string) ([]int, bool, int) {
	var (
		skipNext       bool
		inGarbage      bool
		hasGarbage     bool
		garbageRemoved int
		groupIdx       int
	)

	groups := make([]int, 0)

	for _, char := range input {
		if !skipNext {
			if char == '!' {
				skipNext = true
			}
			if char == '<' && !inGarbage {
				inGarbage = true
				hasGarbage = true
			}
			if char == '>' {
				inGarbage = false
				garbageRemoved--
			}
			if !inGarbage {
				if char == '{' {
					groupIdx++
				}
				if char == '}' {
					groups = append(groups, groupIdx)
					groupIdx--
				}
			}
		} else {
			skipNext = false
		}

		if inGarbage {
			garbageRemoved++
			if skipNext {
				garbageRemoved -= 2
			}
		}
	}

	return groups, hasGarbage, garbageRemoved
}
