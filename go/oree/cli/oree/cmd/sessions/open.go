package sessions

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewOpenCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "open StartTime TrailId StepId",
		Args: cobra.RangeArgs(3, 3),
		Run: func(cmd *cobra.Command, args []string) {
			startTime, useArg := common.TimeArg(args, 0)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[0])
				return
			}
			trailId, _ := common.StringArg(args, 1)
			stepId, _ := common.StringArg(args, 2)
			runWithOree(func(o oree.OreeI) {
				open(o, oree.TrailId(trailId), oree.StepId(stepId), startTime)
			})
		},
	}
	return cmd
}

func open(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, startTime time.Time) {
	trail, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
	}
	step, status := trail.StepWithId(stepId)
	if status == oree.NotFound {
		common.PrintLines(common.FormatIdNotFound("step", stepId))
	}
	o.OpenSessionManager().Open(oree.OpenSession{
		StartTime: startTime,
		Trail:     trail,
		Step:      step,
	})
	list(o)
}
