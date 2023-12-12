package main

type Score []int

func NewScore(rank int, values ...[]int) Score {
	score := make([]int, 0)
	score = append(score, rank)

	for _, val := range values {
		score = append(score, val...)
	}

	for i := len(score); i < 6; i++ {
		score = append(score, 0)
	}

	return score
}

func (s Score) Compare(other Score) int {
	if len(s) != len(other) {
		panic("scores are not equal length")
	}

	for i := 0; i < len(s); i++ {
		if s[i] == other[i] {
			continue
		}

		if s[i] > other[i] {
			return 1
		}

		if s[i] < other[i] {
			return -1
		}
	}

	return 0
}
