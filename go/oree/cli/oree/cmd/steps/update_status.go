package steps

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func newUpdateStatusCmd(name string, status oree.StepStatus, runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  fmt.Sprintf("%v TrailId StepId", name),
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			stepId, _ := common.StringArg(args, 1)
			runWithOree(func(o oree.OreeI) {
				updateStatus(o, oree.TrailId(trailId), oree.StepId(stepId), status)
			})
		},
	}
	return cmd
}

func updateStatus(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, targetStatus oree.StepStatus) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
		return
	}
	stepI, status := trailI.StepWithId(stepId)
	if status == oree.NotFound {
		common.PrintLines(common.FormatIdNotFound("step", stepId))
		return
	}
	if status != targetStatus {
		stepI.UpdateStatus(targetStatus)
	}
	list(o, trailId)
}
