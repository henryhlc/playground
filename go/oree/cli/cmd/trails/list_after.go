package trails

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/common"
	"github.com/spf13/cobra"
)

const defaultListAfterLength = 15

func NewListAfterCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-after TrailId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListAfterLength)

			runWithOree(func(o oree.OreeI) {
				listNAfter(o, oree.TrailId(trailId), n)
			})
		},
	}
	return cmd
}

func listNAfter(o oree.OreeI, trailId oree.TrailId, n int) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		fmt.Printf("No trail found for the given trail id \"%v\".\n", trailId)
		return
	}
	trailIs := o.Trails().NAfter(n, trailI)
	numTrails := len(trailIs)
	total := o.Trails().Len()
	common.PrintLines(
		common.FormatTrails(trailIs),
		[]string{""},
		common.FormatNofM(numTrails, total, "trails"),
	)
}
