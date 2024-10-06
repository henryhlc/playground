package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete TrailId StepId",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			stepId, _ := common.StringArg(args, 1)
			runWithOree(func(o oree.OreeI) {
				delete(o, oree.TrailId(trailId), oree.StepId(stepId))
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("trail", trailId),
		)
		return
	}
	stepI, status := trailI.StepWithId(stepId)
	if status == oree.NotFound {
		common.PrintLines(
			common.FormatIdNotFound("step", stepId),
		)
		return

	}
	trailI.StepsWithStatus(status).Delete(stepI)
	list(o, trailId)
}
