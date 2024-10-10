package blocks

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 30

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "blocks [n]",
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			n, _ := common.IntArgOrDefault(args, 0, defaultListLength)
			runWithOree(func(o oree.OreeI) {
				listN(o, n)
			})
		},
	}
	cmd.AddCommand(NewCreateCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	cmd.AddCommand(NewListBeforeCmd(runWithOree))
	cmd.AddCommand(NewDupWithDateCmd(runWithOree))
	cmd.AddCommand(NewUpdateCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI) {
	listN(o, defaultListLength)
}

func listN(o oree.OreeI, n int) {
	blocks := o.Blocks().LastN(n)
	common.PrintLines(common.ConcatLines(
		common.FormatBlocks(blocks),
		[]string{""},
		common.FormatNofM(len(blocks), o.Blocks().Len(), "blocks"),
	))
}
