package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update TrailId StepId Description",
		Args: cobra.RangeArgs(3, 3),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			stepId, _ := common.StringArg(args, 1)
			description, _ := common.StringArg(args, 2)

			runWithOree(func(o oree.OreeI) {
				update(o, oree.TrailId(trailId), oree.StepId(stepId), description)
			})
		},
	}
	return cmd
}

func update(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, description string) {
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
	stepI.Update(oree.Step{
		Description: description,
	})
	common.PrintLines(
		common.FormatStep(stepI),
	)
}
