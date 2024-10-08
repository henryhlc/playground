package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewMoveBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "move-before TrailId NeighborTrailId",
		Args: cobra.RangeArgs(3, 3),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			trailId, _ := common.StringArg(args, 1)
			nbrTrailId, _ := common.StringArg(args, 2)

			runWithOree(func(o oree.OreeI) {
				moveBefore(o, oree.AreaId(areaId), oree.TrailId(trailId), oree.TrailId(nbrTrailId))
			})
		},
	}
	return cmd
}

func moveBefore(o oree.OreeI, areaId oree.AreaId, trailId oree.TrailId, nbrTrailId oree.TrailId) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("area", areaId))
	}
	trail, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("trail", trailId),
		)
		return
	}
	nbrTrail, ok := o.Trails().WithId(nbrTrailId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("trail", nbrTrailId),
		)
		return
	}
	area.Trails().PlaceBefore(trail, nbrTrail)
	list(o, areaId)
}
