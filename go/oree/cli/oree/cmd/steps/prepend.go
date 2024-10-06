package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewPrependCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "prepend TrailId Description",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			description, _ := common.StringArg(args, 1)
			runWithOree(func(o oree.OreeI) {
				prepend(o, oree.TrailId(trailId), description)
			})
		},
	}
	return cmd
}

func prepend(o oree.OreeI, trailId oree.TrailId, description string) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
		return
	}
	trailI.StepsWithStatus(oree.Active).CreateFront(oree.Step{
		Description: description,
	})
	list(o, trailId)
}
