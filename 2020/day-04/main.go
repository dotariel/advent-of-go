package main

import (
	"dotariel/inputreader"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	valid := 0

	for _, entry := range inputreader.ReadStrings("input.txt", "\n\n") {
		if NewPassport(entry).IsValid() {
			valid++
		}
	}

	println(valid, " valid passports")
}

type Passport struct {
	byr Field
	iyr Field
	eyr Field
	hgt Field
	hcl Field
	ecl Field
	pid Field
	cid Field
}

type Field struct {
	value     string
	validator func(string) bool
}

func NewPassport(s string) Passport {
	p := Passport{}
	parsed := make(map[string]string)

	for _, field := range strings.Fields(s) {
		parts := strings.Split(field, ":")
		key := parts[0]
		val := parts[1]

		parsed[key] = val
	}

	p.byr = Field{getField(parsed, "byr"), byr}
	p.iyr = Field{getField(parsed, "iyr"), iyr}
	p.eyr = Field{getField(parsed, "eyr"), eyr}
	p.hgt = Field{getField(parsed, "hgt"), hgt}
	p.hcl = Field{getField(parsed, "hcl"), hcl}
	p.ecl = Field{getField(parsed, "ecl"), ecl}
	p.pid = Field{getField(parsed, "pid"), pid}
	p.cid = Field{getField(parsed, "cid"), cid}

	return p
}

func (f Field) IsValid() bool {
	return f.validator(f.value)
}

func (p Passport) IsValid() bool {
	return p.byr.IsValid() &&
		p.iyr.IsValid() &&
		p.eyr.IsValid() &&
		p.hgt.IsValid() &&
		p.hcl.IsValid() &&
		p.ecl.IsValid() &&
		p.pid.IsValid() &&
		p.cid.IsValid()
}

func byr(s string) bool {
	return intRange(s, 1920, 2002)
}

func iyr(s string) bool {
	return intRange(s, 2010, 2020)
}

func eyr(s string) bool {
	return intRange(s, 2020, 2030)
}

func hgt(s string) bool {
	if strings.HasSuffix(s, "cm") {
		return intRange(strings.TrimRight(s, "cm"), 150, 193)
	}

	if strings.HasSuffix(s, "in") {
		return intRange(strings.TrimRight(s, "in"), 59, 76)
	}

	return false
}

func hcl(s string) bool {
	if !strings.HasPrefix(s, "#") {
		return false
	}

	hex := s[1:]

	if len(hex) != 6 {
		return false
	}

	for i := 0; i < len(hex); i++ {
		if !unicode.Is(unicode.ASCII_Hex_Digit, rune(hex[i])) {
			return false
		}
	}

	return true
}

func ecl(s string) bool {
	for _, val := range strings.Fields("amb blu brn gry grn hzl oth") {
		if val == s {
			return true
		}
	}

	return false
}

func pid(s string) bool {
	if len(s) != 9 {
		return false
	}

	for _, val := range s {
		if !intRange(string(val), 0, 9) {
			return false
		}
	}

	return true
}

func cid(s string) bool {
	return true
}

func intRange(s string, min int, max int) bool {
	i, err := strconv.Atoi((s))

	if err != nil {
		return false
	}

	return i >= min && i <= max
}

func getField(m map[string]string, key string) string {
	if val, ok := m[key]; ok {
		return val
	}

	return ""
}
