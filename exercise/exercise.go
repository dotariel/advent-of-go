package exercise

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dotariel/advent-of-go/library/captcha"
	"github.com/dotariel/advent-of-go/library/checksum"
	"github.com/dotariel/advent-of-go/library/memory"
	"github.com/dotariel/advent-of-go/library/passphrase"
	"github.com/dotariel/advent-of-go/library/register"
	"github.com/dotariel/advent-of-go/library/stack"
	"github.com/dotariel/advent-of-go/library/stream"
	"github.com/dotariel/advent-of-go/library/tower"
)

var exercises = make(map[int]Exercise)

type Part func(string) interface{}
type Exercise []Part

func init() {
	exercises[1] = Exercise([]Part{
		func(input string) interface{} { return captcha.InverseCaptcha(input, captcha.WrapAroundAdd) },
		func(input string) interface{} { return captcha.InverseCaptcha(input, captcha.HalfwayAroundAdd) },
	})
	exercises[2] = Exercise([]Part{
		func(input string) interface{} { return checksum.NewSpreadsheet(input).Checksum(checksum.Diff) },
		func(input string) interface{} { return checksum.NewSpreadsheet(input).Checksum(checksum.Factor) },
	})
	exercises[3] = Exercise([]Part{
		func(input string) interface{} { return "NOT IMPLEMENTED" },
		func(input string) interface{} { return "NOT IMPLEMENTED" },
	})
	exercises[4] = Exercise([]Part{
		func(input string) interface{} { return passphrase.CountValid(input, passphrase.Unique) },
		func(input string) interface{} { return passphrase.CountValid(input, passphrase.NonAnagram) },
	})
	exercises[5] = Exercise([]Part{
		func(input string) interface{} { return stack.NewStack(input).Trace(stack.SimpleIncrementer) },
		func(input string) interface{} { return stack.NewStack(input).Trace(stack.BiasedDecrementer) },
	})
	exercises[6] = Exercise([]Part{
		func(input string) interface{} { return memory.NewState(input).CountCycles() },
		// func(input string) interface{} { return stack.NewStack(input).Trace(stack.BiasedDecrementer) },
	})
	exercises[7] = Exercise([]Part{
		func(input string) interface{} { return tower.New(input).Name },
		func(input string) interface{} {
			return tower.FindMismatch(tower.New(input))
		},
	})
	exercises[8] = Exercise([]Part{
		func(input string) interface{} {
			registers := register.New()
			registers.ProcessBatch(input)

			return registers.Max()
		},
		func(input string) interface{} {
			registers := register.New()
			return registers.ProcessBatch(input)
		},
	})
	exercises[9] = Exercise([]Part{
		func(input string) interface{} {
			score := 0
			groups, _, _ := stream.Parse(input)
			for _, group := range groups {
				score += group
			}

			return score
		},
		func(input string) interface{} {
			_, _, removed := stream.Parse(input)
			return removed
		},
	})

}

func Run(day int) error {
	file := fmt.Sprintf("%v/src/github.com/dotariel/advent-of-go/exercise/inputs/%d", os.Getenv("GOPATH"), day)
	bytes, _ := ioutil.ReadFile(file)

	exercise, ok := exercises[day]

	if !ok {
		return fmt.Errorf("no exercise found for day %v", day)
	}

	for part, fn := range exercise {
		fmt.Printf("Part %v: %v\n", part+1, fn(string(bytes)))
	}

	return nil
}
