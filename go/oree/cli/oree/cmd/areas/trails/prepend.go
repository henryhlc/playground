package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewPrependCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "prepend AreaId TrailId",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			trailIds := []oree.TrailId{}
			for _, trailId := range args[1:] {
				trailIds = append(trailIds, oree.TrailId(trailId))
			}
			runWithOree(func(o oree.OreeI) {
				prepend(o, oree.AreaId(areaId), trailIds)
			})
		},
	}
	return cmd
}

func prepend(o oree.OreeI, areaId oree.AreaId, trailIds []oree.TrailId) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("area", areaId))
		return
	}
	trails := []oree.TrailI{}
	for _, trailId := range trailIds {
		trail, ok := o.Trails().WithId(trailId)
		if !ok {
			common.PrintLines(common.FormatIdNotFound("trail", trailId))
			return
		}
		trails = append(trails, trail)
	}
	for i := len(trails) - 1; i >= 0; i-- {
		area.Trails().PlaceFront(trails[i])
	}
	list(o, areaId)
}
