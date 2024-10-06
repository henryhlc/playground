package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/spf13/cobra"
)

func NewActivateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	return newUpdateStatusCmd("activate", oree.Active, runWithOree)

}
