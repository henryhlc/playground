package areas

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultTrailsListLength = 15

func NewPrependCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "prepend Description [t1 t2 ...]",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			description, _ := common.StringArg(args, 0)
			trailIds := []oree.TrailId{}
			for _, trailId := range args[1:] {
				trailIds = append(trailIds, oree.TrailId(trailId))
			}
			runWithOree(func(o oree.OreeI) {
				prepend(o, description, trailIds)
			})
		},
	}
	return cmd
}

func prepend(o oree.OreeI, description string, trailIds []oree.TrailId) {
	trails := []oree.TrailI{}
	for _, trailId := range trailIds {
		trail, ok := o.Trails().WithId(trailId)
		if !ok {
			common.PrintLines(common.FormatIdNotFound("trail", trailId))
			return
		}
		trails = append(trails, trail)
	}

	area := o.Areas().CreateFront(oree.Area{
		Description: description,
	})
	for _, trail := range trails {
		area.Trails().PlaceBack(trail)
	}
	common.PrintLines(
		common.FormatAreaWithTrails(area, defaultTrailsListLength),
	)
}
