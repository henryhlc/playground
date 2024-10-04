package common

import "fmt"

func IntArgOrDefault(args []string, idx int, d int) (n int, useArg bool) {
	if idx >= len(args) {
		return d, false
	}
	var argN int
	_, err := fmt.Sscan(args[idx], &argN)
	if err != nil {
		return d, false
	}
	return argN, true
}

func StringArg(args []string, idx int) (s string, useArg bool) {
	if idx >= len(args) {
		return "", false
	}
	return args[idx], true
}
