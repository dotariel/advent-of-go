package register

import (
	"reflect"
	"testing"
)

func TestOperations(t *testing.T) {
	testCases := []struct {
		a         int
		b         int
		operation string
		expected  int
	}{
		{2, 3, "inc", 5},
		{0, 0, "dec", 0},
		{0, 0, "inc", 0},
		{4, 5, "dec", -1},
		{1, 1, "noop", 1},
	}

	for _, tt := range testCases {
		if actual := operationFromString(tt.operation)(tt.a, tt.b); actual != tt.expected {
			t.Errorf("failed %v(%v,%v); wanted:%v but got:%v", tt.operation, tt.a, tt.b, tt.expected, actual)
		}
	}
}

func TestConditions(t *testing.T) {
	testCases := []struct {
		a         int
		b         int
		condition string
		expected  bool
	}{
		{5, 5, "==", true},
		{4, 5, "==", false},
		{6, 5, ">=", true},
		{5, 5, ">=", true},
		{4, 5, ">=", false},
		{4, 5, "<=", true},
		{4, 3, "<=", false},
		{4, 4, "<=", true},
		{0, 1, "!=", true},
	}

	for _, tt := range testCases {
		if actual := conditionFromString(tt.condition)(tt.a, tt.b); actual != tt.expected {
			t.Errorf("failed %v(%v,%v); wanted:%v but got:%v", tt.condition, tt.a, tt.b, tt.expected, actual)
		}
	}
}

func TestProcess(t *testing.T) {
	testCases := []struct {
		input     string
		registers Registers
		expected  Registers
	}{
		{"b inc 5 if a > 1", New(), map[string]int{"a": 0, "b": 0}},
		{"a inc 1 if b < 5", map[string]int{"a": 0, "b": 0}, map[string]int{"a": 1, "b": 0}},
		{"c dec -10 if a >= 1", map[string]int{"a": 1, "b": 0}, map[string]int{"a": 1, "b": 0, "c": 10}},
		{"c inc -20 if c == 10", map[string]int{"a": 1, "b": 0, "c": 10}, map[string]int{"a": 1, "b": 0, "c": -10}},
		{"foo inc 1 if bar == 0", New(), map[string]int{"foo": 1, "bar": 0}},
		{"foo dec 143 if bar == 0", New(), map[string]int{"foo": -143, "bar": 0}},
		{"l inc 669 if d != -754", New(), map[string]int{"l": 669, "d": 0}},
	}

	for _, tt := range testCases {
		if tt.registers.Process(tt.input); !reflect.DeepEqual(tt.registers, tt.expected) {
			t.Errorf("Failed %v; wanted:%v, but got:%v", tt.input, tt.expected, tt.registers)
		}
	}
}

func TestRegisters_Max(t *testing.T) {
	testCases := []struct {
		name     string
		input    Registers
		expected int
	}{
		{"all zeros", Registers(map[string]int{"a": 0, "b": 0, "c": 0}), 0},
		{"some positive, some negative", Registers(map[string]int{"a": -14, "b": 0, "c": 27}), 27},
		{"all positive", Registers(map[string]int{"a": 0, "b": 1, "c": 3, "d": 5, "e": 7, "f": 9}), 9},
		{"all negative", Registers(map[string]int{"a": -1, "b": -5, "c": -14}), -1},
	}

	for _, tt := range testCases {
		if actual := tt.input.Max(); actual != tt.expected {
			t.Errorf("failed assertion for '%v (%v)'; wanted:%v, but got:%v", tt.name, tt.input, tt.expected, actual)
		}
	}
}

func TestRegisters_ProcessBatch(t *testing.T) {
	input := `b inc 5 if a > 1
            a inc 1 if b < 5
            c dec -10 if a >= 1
            c inc -20 if c == 10`

	r := New()
	highest := r.ProcessBatch(input)

	expected := struct{ max, high int }{1, 10}

	if r.Max() != expected.max {
		t.Errorf("failed; wanted:%v, but got:%v", expected.max, r.Max())
	}

	if highest != expected.high {
		t.Errorf("failed; wanted:%v, but got:%v", expected.high, highest)
	}
}
