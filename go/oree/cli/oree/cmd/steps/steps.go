package steps

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 5

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "steps TrailId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListLength)
			runWithOree(func(o oree.OreeI) {
				listN(o, oree.TrailId(trailId), n)
			})
		},
	}
	cmd.AddCommand(NewPrependCmd(runWithOree))
	cmd.AddCommand(NewActivateCmd(runWithOree))
	cmd.AddCommand(NewPinCmd(runWithOree))
	cmd.AddCommand(NewArchiveCmd(runWithOree))
	cmd.AddCommand(NewListAfterCmd(runWithOree))
	cmd.AddCommand(NewUpdateCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	cmd.AddCommand(NewMoveBeforeCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI, trailId oree.TrailId) {
	listN(o, trailId, defaultListLength)
}

func listN(o oree.OreeI, trailId oree.TrailId, n int) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
		return
	}
	common.PrintLines(
		common.FormatTrailWithSteps(trailI, []oree.StepStatus{oree.Pinned, oree.Active, oree.Archived}, n),
	)
}
