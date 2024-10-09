package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 15

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "trails AreaId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListLength)
			runWithOree(func(o oree.OreeI) {
				listN(o, oree.AreaId(areaId), n)
			})
		},
	}
	cmd.AddCommand(NewListAfterCmd(runWithOree))
	cmd.AddCommand(NewPrependCmd(runWithOree))
	cmd.AddCommand(NewMoveBeforeCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI, areaId oree.AreaId) {
	listN(o, areaId, defaultListLength)
}

func listN(o oree.OreeI, areaId oree.AreaId, n int) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("area", areaId),
		)
		return
	}
	trails := area.Trails().FirstN(n)
	numTrails := len(trails)
	total := area.Trails().Len()
	common.PrintLines(
		common.FormatArea(area),
		common.FormatPrefix("  ", common.FormatTrails(trails)),
		[]string{""},
		common.FormatNofM(numTrails, total, "trails"),
	)
}
