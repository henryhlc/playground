package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewMoveBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "move-before TrailId StepId NeighborStepId",
		Args: cobra.RangeArgs(3, 3),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			stepId, _ := common.StringArg(args, 1)
			nbrStepId, _ := common.StringArg(args, 2)
			runWithOree(func(o oree.OreeI) {
				moveBefore(o, oree.TrailId(trailId), oree.StepId(stepId), oree.StepId(nbrStepId))
			})
		},
	}
	return cmd
}

func moveBefore(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, nbrStepId oree.StepId) {
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
	nbrStepI, nbrStatus := trailI.StepWithId(nbrStepId)
	if nbrStatus == oree.NotFound {
		common.PrintLines(
			common.FormatIdNotFound("step", nbrStepId),
		)
		return
	}
	stepI.UpdateStatus(nbrStatus)
	trailI.StepsWithStatus(nbrStatus).PlaceBefore(stepI, nbrStepI)
	list(o, trailId)
}
