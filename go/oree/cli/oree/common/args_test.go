package common_test

import (
	"testing"

	"github.com/henryhlc/playground/go/oree/cli/common"
)

func TestIntArgOrDef(t *testing.T) {
	tests := []struct {
		scenario       string
		args           []string
		idx            int
		d              int
		expectedN      int
		expectedUseArg bool
	}{
		{"Out of range", []string{"12"}, 1, 3, 3, false},
		{"Not a number", []string{"abc"}, 0, 3, 3, false},
		{"A number", []string{"abc", "2"}, 1, 3, 2, true},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			actualN, actualUseArg := common.IntArgOrDefault(test.args, test.idx, test.d)
			if actualN != test.expectedN {
				t.Errorf("Expected n: %v, actual: %v", test.expectedN, actualN)
			}
			if actualUseArg != test.expectedUseArg {
				t.Errorf("Expected useArg: %v, actual: %v", test.expectedUseArg, actualUseArg)
			}
		})
	}
}

func TestStringArg(t *testing.T) {
	tests := []struct {
		scenario       string
		args           []string
		idx            int
		expectedS      string
		expectedUseArg bool
	}{
		{"Out of range", []string{}, 0, "", false},
		{"A string", []string{"abc", "def"}, 1, "def", true},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			actualS, actualUseArg := common.StringArg(test.args, test.idx)
			if actualS != test.expectedS {
				t.Errorf("Expected s: %v, actual: %v", test.expectedS, actualS)
			}
			if actualUseArg != test.expectedUseArg {
				t.Errorf("Expected useArg: %v, actual: %v", test.expectedUseArg, actualUseArg)
			}
		})
	}
}
