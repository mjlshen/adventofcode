package main

import "testing"

func TestGammaEpsilon(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 198,
		},
	}

	for _, test := range tests {
		actual := gammaEpsilon(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestO2CO2(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected int
	}{
		{
			input:    "test.txt",
			length:   5,
			expected: 230,
		},
	}

	for _, test := range tests {
		actual := o2CO2(test.input, test.length)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func Test_ReadingsInsert(t *testing.T) {
	tests := []struct {
		input               []string
		length              int
		expectedOnesCount   int
		expectedZeroesCount int
	}{
		{
			input:               []string{"0", "1"},
			length:              1,
			expectedZeroesCount: 1,
			expectedOnesCount:   1,
		},
		{
			input:               []string{"0001", "0010", "1100"},
			length:              4,
			expectedZeroesCount: 2,
			expectedOnesCount:   1,
		},
	}

	for _, test := range tests {
		readings := &Readings{
			SortDepth: 0,
			Length:    test.length,
		}
		for _, input := range test.input {
			readings.insert(input)
		}
		if len(readings.Ones) != test.expectedOnesCount {
			t.Errorf("Expected %d ones, got %d", test.expectedOnesCount, len(readings.Ones))
		} else if len(readings.Zeros) != test.expectedZeroesCount {
			t.Errorf("Expected %d zeroes, got %d", test.expectedZeroesCount, len(readings.Zeros))
		}
	}
}

func Test_ReadingsO2(t *testing.T) {
	tests := []struct {
		input    []string
		length   int
		expected string
	}{
		{
			input:    []string{},
			length:   0,
			expected: "",
		},
		{
			input:    []string{"0001", "1000", "0010", "1100", "0110"},
			length:   4,
			expected: "0010",
		},
	}

	for _, test := range tests {
		readings := &Readings{
			SortDepth: 0,
			Length:    test.length,
		}
		for _, input := range test.input {
			readings.insert(input)
		}
		actual := readings.O2()
		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
		}
	}
}

func Test_ReadingsCO2(t *testing.T) {
	tests := []struct {
		input    []string
		length   int
		expected string
	}{
		{
			input:    []string{},
			length:   0,
			expected: "",
		},
		{
			input:    []string{"0001", "1000", "0010", "1100", "0110"},
			length:   4,
			expected: "1000",
		},
	}

	for _, test := range tests {
		readings := &Readings{
			SortDepth: 0,
			Length:    test.length,
		}
		for _, input := range test.input {
			readings.insert(input)
		}
		actual := readings.CO2()
		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
		}
	}
}
