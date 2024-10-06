package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/spf13/cobra"
)

func NewArchiveCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	return newUpdateStatusCmd("archive", oree.Archived, runWithOree)

}
