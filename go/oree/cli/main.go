package main

import (
	"fmt"
	"os"

	"github.com/henryhlc/playground/go/oree/cli/cmd"
)

func main() {
	if err := cmd.NewCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
