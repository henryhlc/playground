package oreejson

import "testing"

func TestGetAndIncId(t *testing.T) {
	oj := FromData(NewOreeJD())
	ids := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "aa", "ba", "ca",
		"da", "ea", "fa", "ga", "ha",
	}
	for i := range ids {
		actual := oj.getAndIncId()
		if expected := ids[i]; actual != expected {
			t.Errorf("Id expected %v, actual %v", expected, actual)
		}
	}
}
