package main

import (
	"dotariel/inputreader"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewField(t *testing.T) {
	field := NewField(file("simple.txt"))

	assert.Len(t, field, 5)
	assert.Len(t, [][]*Tile(field)[0], 5)

}

func TestField_Get(t *testing.T) {
	field := NewField(file("simple.txt"))

	field.Get(Point{0, 0})
}

func TestField_Walk(t *testing.T) {
	field := NewField(file("complex.txt"))

	path := field.Walk()

	assert.Equal(t, 8, path.Max())
}

func TestField_NextFrom(t *testing.T) {
	field := NewField(file("simple.txt"))

	start, _ := field.Start()
	connector, nextDirection := field.Next(start, SOUTH)

	assert.NotNil(t, connector)
	assert.Equal(t, SOUTH, nextDirection)
}

func TestOther(t *testing.T) {
	field := NewField(inputreader.ReadAll("input.txt"))

	fmt.Println(getInteriorPoints(field))
}

func file(filename string) string {
	return inputreader.ReadAll("samples/" + filename)
}

func getInteriorPoints(f Field) int {
	path := f.Walk()

	inner := path.Area() - (len(path) / 2) + 1

	return inner
}
