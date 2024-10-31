package main

import (
	"math"
	"testing"
)

// Tolerance for comparing floating-point results
const epsilon = 0.000001

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}

// Table-driven tests
var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.dividend, tt.divisor)
			if tt.isErr {
				if err == nil {
					t.Error("Expected an error but did not get one")
				}
			} else {
				if err != nil {
					t.Error("Did not expect an error, but got one:", err)
				} else if !floatEquals(got, tt.expected) {
					t.Errorf("Expected %f but got %f", tt.expected, got)
				}
			}
		})
	}
}
