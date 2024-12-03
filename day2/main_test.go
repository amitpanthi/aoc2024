package main

import "testing"

type testCase struct {
	report   string
	expected bool
}

func TestIsValid(t *testing.T) {
	tests := []testCase{
		{
			report:   "1 3 5 7 11 9",
			expected: true,
		}, {
			report:   "1 3 5 7 11 11",
			expected: false,
		}, {
			report:   "11 9 5 7 5",
			expected: true,
		}, {
			report:   "7 6 4 2 1",
			expected: true,
		}, {
			report:   "1 2 7 8 9",
			expected: false,
		}, {
			report:   "9 7 6 2 1",
			expected: false,
		}, {
			report:   "1 3 2 4 5",
			expected: true,
		}, {
			report:   "8 6 4 4 1",
			expected: true,
		}, {
			report:   "1 3 6 7 9",
			expected: true,
		}, {
			report:   "7 10 8 10 11",
			expected: true,
		}, {
			report:   "29 28 27 25 26 25 22 20",
			expected: true,
		}, {
			report:   "9 8 7 6 7",
			expected: true,
		}, {
			report:   "1 2 3 4 3",
			expected: true,
		}, {
			report:   "1 6 7 8 9",
			expected: true,
		}, {
			report:   "1 4 3 2 1",
			expected: true,
		}, {
			report:   "5 1 2 3 4 5",
			expected: true,
		}, {
			report:   "1 2 3 4 5 5",
			expected: true,
		}, {
			report:   "1 1 2 3 4 5",
			expected: true,
		}, {
			report:   "48 46 47 49 51 54 56",
			expected: true,
		},
	}

	for _, test := range tests {
		valid := isValidReport(test.report)

		if valid != test.expected {
			t.Errorf("Test failed for report string %s, expected %t, got %t",
				test.report, test.expected, valid)
		}
	}
}
