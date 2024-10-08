package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete AreaId TrailId",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			trailId, _ := common.StringArg(args, 1)

			runWithOree(func(o oree.OreeI) {
				delete(o, oree.AreaId(areaId), oree.TrailId(trailId))
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, areaId oree.AreaId, trailId oree.TrailId) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("area", areaId))
		return
	}
	trail, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
		return
	}
	area.Trails().Delete(trail)
	list(o, areaId)
}
