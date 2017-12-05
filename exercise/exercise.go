package exercise

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dotariel/advent-of-go/library/captcha"
	"github.com/dotariel/advent-of-go/library/checksum"
	"github.com/dotariel/advent-of-go/library/passphrase"
	"github.com/dotariel/advent-of-go/library/stack"
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
}

func Run(day int) {
	file := fmt.Sprintf("%v/src/github.com/dotariel/advent-of-go/exercise/inputs/%d", os.Getenv("GOPATH"), day)
	bytes, _ := ioutil.ReadFile(file)

	exercise := exercises[day]

	fmt.Println("Part 1: ", exercise[0](string(bytes)))
	fmt.Println("Part 2: ", exercise[1](string(bytes)))
}
