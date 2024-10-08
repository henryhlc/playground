package areas

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewMoveBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "move-before AreaId NeighborAreaId",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			nbrAreaId, _ := common.StringArg(args, 1)

			runWithOree(func(o oree.OreeI) {
				moveBefore(o, oree.AreaId(areaId), oree.AreaId(nbrAreaId))
			})
		},
	}
	return cmd
}

func moveBefore(o oree.OreeI, areaId oree.AreaId, nbrAreaId oree.AreaId) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("area", areaId))
		return
	}
	nbrArea, ok := o.Areas().WithId(nbrAreaId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("area", nbrAreaId))
		return
	}
	o.Areas().PlaceBefore(area, nbrArea)
	list(o)
}
