package main

import (
	"os"
	"testing"
)

func TestNumBeacons(t *testing.T) {
	tests := []struct {
		input           string
		expectedBeacons int
		expectedMaxDist int
	}{
		{
			input:           "test.txt",
			expectedBeacons: 79,
			expectedMaxDist: 3621,
		},
	}

	for _, test := range tests {
		actualBeacons, actualMaxDist := numBeacons(test.input)
		if actualBeacons != test.expectedBeacons {
			t.Errorf("actual == %d, expected %d", actualBeacons, test.expectedBeacons)
		}
		if actualMaxDist != test.expectedMaxDist {
			t.Errorf("actual == %d, expected %d", actualMaxDist, test.expectedMaxDist)
		}
	}
}

func TestGetRotations(t *testing.T) {
	if len(getRotations()) != 24 {
		t.Errorf("actual == %d, expected %d", len(getRotations()), 24)
	}
}

func Test_Mat3DDeterminant(t *testing.T) {
	tests := []struct {
		mat      Mat3D
		expected int
	}{
		{
			mat: Mat3D{
				{6, 1, 1},
				{4, -2, 5},
				{2, 8, 7},
			},
			expected: -306,
		},
	}

	for _, test := range tests {
		actual := test.mat.determinant()
		if actual != test.expected {
			t.Errorf("actual == %d, expected %d", actual, test.expected)
		}
	}
}

func Test_ScannersString(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: "test.txt",
		},
	}

	for _, test := range tests {
		actual := parseScanners(test.input).String()
		expected, _ := os.ReadFile(test.input)
		if actual != string(expected) {
			t.Errorf("actual == %s, expected %s", actual, string(expected))
		}
	}
}
