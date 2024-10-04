package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 15

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "trails [n]",
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			n, _ := common.IntArgOrDefault(args, 0, defaultListLength)
			runWithOree(func(o oree.OreeI) {
				listN(o, n)
			})
		},
	}
	cmd.AddCommand(NewListAfterCmd(runWithOree))
	cmd.AddCommand(NewPrependCmd(runWithOree))
	cmd.AddCommand(NewUpdateCmd(runWithOree))
	cmd.AddCommand(NewMoveBeforeCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI) {
	listN(o, defaultListLength)
}

func listN(o oree.OreeI, n int) {
	trailIs := o.Trails().FirstN(n)
	numTrails := len(trailIs)
	total := o.Trails().Len()
	common.PrintLines(
		common.FormatTrails(trailIs),
		[]string{""},
		common.FormatNofM(numTrails, total, "trails"),
	)
}
