package algo

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		description  string
		inputNums    []int
		expectedLow  int
		expectedHigh int
	}{
		{"EmptyNums", []int{}, -1, 0},
		{"OneTrue", []int{1}, -1, 0},
		{"OneFalse", []int{0}, 0, 1},
		{"TwoAll", []int{1, 1}, -1, 0},
		{"TwoOne", []int{0, 1}, 0, 1},
		{"TwoNone", []int{0, 0}, 1, 2},
		{"ThreeAll", []int{1, 1, 1}, -1, 0},
		{"ThreeTwo", []int{0, 1, 1}, 0, 1},
		{"ThreeOne", []int{0, 0, 1}, 1, 2},
		{"ThreeNone", []int{0, 0, 0}, 2, 3},
		{"FourAll", []int{1, 1, 1, 1}, -1, 0},
		{"FourThree", []int{0, 1, 1, 1}, 0, 1},
		{"FourTwo", []int{0, 0, 1, 1}, 1, 2},
		{"FourOne", []int{0, 0, 0, 1}, 2, 3},
		{"FourNone", []int{0, 0, 0, 0}, 3, 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			l, h := BinarySearch(testCase.inputNums, func(n int) bool { return n > 0 })
			if l != testCase.expectedLow {
				t.Errorf("expectedLow=%v; actualLow=%v", testCase.expectedLow, l)
			}
			if h != testCase.expectedHigh {
				t.Errorf("expectedHigh=%v; actualHigh=%v", testCase.expectedHigh, h)
			}
		})
	}
}
