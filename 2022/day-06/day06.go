package main

func FindMarker(s string) int {
	return FindN(s, 4)
}

func FindMessage(s string) int {
	return FindN(s, 14)
}

func FindN(s string, offset int) int {
	for i := 0; i < len(s)-(offset-1); i++ {
		buffer := []rune(s[i : i+offset])
		unique := Unique(buffer)
		found := len(buffer) != len(unique)

		if !found {
			return offset + i
		}
	}

	return -1
}

func Unique(input []rune) []rune {
	keys := make(map[rune]bool)
	uniq := []rune{}

	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniq = append(uniq, entry)
		}
	}

	return uniq
}
