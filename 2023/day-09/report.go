package main

import (
	"dotariel/util"
	"strings"
)

type Report []History

type History []int

type Sequence []int

func NewReport(input string) Report {
	report := Report{}

	for _, line := range strings.Split(input, "\n") {
		report = append(report, util.ToInts(strings.Fields(line)))
	}

	return report
}

func (h History) PredictNext() int {
	sequences := make([]Sequence, 0)

	seq := Sequence(h)
	sequences = append(sequences, seq)

	for {
		next := seq.Next()

		sequences = append(sequences, next)
		if next.AllZeros() {
			break
		}

		seq = next
	}

	for i := len(sequences) - 1; i >= 0; i-- {
		currentseq := sequences[i]

		if i < len(sequences) && i > 0 {
			targetSeq := sequences[i-1]
			left := targetSeq[len(targetSeq)-1]
			below := currentseq[len(currentseq)-1]
			prediction := left + below

			sequences[i-1] = append(sequences[i-1], prediction)
		}
	}

	top := sequences[0]

	return top[len(top)-1]
}

func (h History) PredictPrevious() int {
	sequences := make([]Sequence, 0)

	seq := Sequence(h)
	sequences = append(sequences, seq)

	for {
		next := seq.Next()

		sequences = append(sequences, next)
		if next.AllZeros() {
			break
		}

		seq = next
	}

	for i := len(sequences) - 1; i >= 0; i-- {
		currentseq := sequences[i]

		if i < len(sequences) && i > 0 {
			targetSeq := sequences[i-1]
			left := targetSeq[0]
			below := currentseq[0]
			prediction := left - below

			sequences[i-1] = append([]int{prediction}, sequences[i-1]...)
		}
	}

	top := sequences[0]

	return top[0]
}

func (s Sequence) Next() Sequence {
	seq := make(Sequence, 0)

	for i := 0; i < len(s)-1; i++ {
		a, b := i, i+1

		seq = append(seq, s[b]-s[a])
	}

	return seq
}

func (s Sequence) AllZeros() bool {
	allzeros := true

	for _, val := range s {
		allzeros = allzeros && val == 0
	}

	return allzeros
}
