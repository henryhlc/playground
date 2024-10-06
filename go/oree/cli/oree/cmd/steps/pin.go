package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/spf13/cobra"
)

func NewPinCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	return newUpdateStatusCmd("pin", oree.Pinned, runWithOree)

}
