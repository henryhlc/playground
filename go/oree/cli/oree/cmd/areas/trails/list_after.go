package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListAfterLength = 15

func NewListAfterCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-after AreaId TrailId [n]",
		Args: cobra.RangeArgs(2, 3),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			trailId, _ := common.StringArg(args, 1)
			n, _ := common.IntArgOrDefault(args, 2, defaultListAfterLength)

			runWithOree(func(o oree.OreeI) {
				listNAfter(o, oree.AreaId(areaId), oree.TrailId(trailId), n)
			})
		},
	}
	return cmd
}

func listNAfter(o oree.OreeI, areaId oree.AreaId, trailId oree.TrailId, n int) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("area", areaId),
		)
		return

	}
	trail, ok := area.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("trail", trailId),
		)
		return
	}

	trails := area.Trails().NAfter(n, trail)
	numTrails := len(trails)
	total := area.Trails().Len()
	common.PrintLines(
		common.FormatArea(area),
		common.FormatPrefix("  ", common.FormatTrails(trails)),
		[]string{""},
		common.FormatNofM(numTrails, total, "trails"),
	)
}
