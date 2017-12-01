package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dotariel/advent-of-go/library/captcha"
)

func main() {
	bytes, _ := ioutil.ReadFile("input")
	input := string(bytes)

	fmt.Println("Part 1: ", captcha.InverseCaptcha(input, captcha.WrapAroundAdd))
	fmt.Println("Part 2: ", captcha.InverseCaptcha(input, captcha.HalfwayAroundAdd))
}
