package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListAfterLength = 15

func NewListAfterCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-after TrailId StepId [n]",
		Args: cobra.RangeArgs(2, 3),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			stepId, _ := common.StringArg(args, 1)
			n, _ := common.IntArgOrDefault(args, 2, defaultListAfterLength)

			runWithOree(func(o oree.OreeI) {
				listNAfter(o, oree.TrailId(trailId), oree.StepId(stepId), n)
			})
		},
	}
	return cmd
}

func listNAfter(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, n int) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("trail", trailId),
		)
		return
	}
	stepI, status := trailI.StepWithId(stepId)
	stepsI := trailI.StepsWithStatus(status)
	common.PrintLines(
		common.FormatTrail(trailI),
		common.FormatStepsSection(status, stepsI.Len(), stepsI.NAfter(n, stepI)),
	)
}
