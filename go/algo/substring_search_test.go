package algo

import (
	"slices"
	"testing"
)

func TestFirstNOccurancesr(t *testing.T) {
	testCases := []struct {
		description string
		s, substr   string
		n           int
		expected    []int
	}{
		{"basic", "abcde", "bc", 1, []int{1}},
		{"all", "asdasdfeasdf", "asdf", 0, []int{3, 8}},
		{"limit", "asdasdfeasdf", "asdf", 1, []int{3}},
		{"overlap", "abcabcabc", "abcabc", 3, []int{0, 3}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := FirstNSubstrOccurrances(testCase.s, testCase.substr, testCase.n)
			if !slices.Equal(actual, testCase.expected) {
				t.Errorf("case:%v, s:%v, substr:%v, expected:%v, actual:%v",
					testCase.description,
					testCase.s,
					testCase.substr,
					testCase.expected,
					actual)
			}
		})
	}
}

func TestFirstOccurrance(t *testing.T) {
	testCases := []struct {
		description string
		s, substr   string
		expected    int
	}{
		{"basic", "abcde", "bc", 1},
		{"multiple", "asdasdfeasdf", "asdf", 3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := FirstSubstrOccurrance(testCase.s, testCase.substr)
			if actual != testCase.expected {
				t.Errorf("case:%v, s:%v, substr:%v, expected:%v, actual:%v",
					testCase.description,
					testCase.s,
					testCase.substr,
					testCase.expected,
					actual)
			}
		})
	}
}

func TestAllOccurrances(t *testing.T) {
	testCases := []struct {
		description string
		s, substr   string
		expected    []int
	}{
		{"basic", "abcde", "bc", []int{1}},
		{"multiple", "asdasdfeasdf", "asdf", []int{3, 8}},
		{"overlap", "abcabcabc", "abcabc", []int{0, 3}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := AllSubstrOccurrances(testCase.s, testCase.substr)
			if !slices.Equal(actual, testCase.expected) {
				t.Errorf("case:%v, s:%v, substr:%v, expected:%v, actual:%v",
					testCase.description,
					testCase.s,
					testCase.substr,
					testCase.expected,
					actual)
			}
		})
	}
}
