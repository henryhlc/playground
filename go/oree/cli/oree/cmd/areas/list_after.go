package areas

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListAfterLength = 15

func NewListAfterCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-after AreaId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListAfterLength)

			runWithOree(func(o oree.OreeI) {
				listNAfter(o, oree.AreaId(areaId), n)
			})
		},
	}
	return cmd
}

func listNAfter(o oree.OreeI, areaId oree.AreaId, n int) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("area", areaId),
		)
		return
	}
	areas := o.Areas().NAfter(n, area)
	numAreas := len(areas)
	total := o.Areas().Len()
	common.PrintLines(
		common.FormatAreas(areas),
		[]string{""},
		common.FormatNofM(numAreas, total, "areas"),
	)
}
