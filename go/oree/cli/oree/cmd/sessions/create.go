package sessions

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewCreateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create StartTime EndTime TrailId StepId",
		Args: cobra.RangeArgs(4, 4),
		Run: func(cmd *cobra.Command, args []string) {
			startTime, useArg := common.TimeArg(args, 0)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[0])
				return
			}
			endTime, useArg := common.TimeArg(args, 1)
			if !useArg {
				fmt.Printf("Invalid duration %v\n", args[1])
			}
			trailId, _ := common.StringArg(args, 2)
			stepId, _ := common.StringArg(args, 3)
			runWithOree(func(o oree.OreeI) {
				create(o, oree.TrailId(trailId), oree.StepId(stepId), startTime, endTime)
			})
		},
	}
	return cmd
}

func create(o oree.OreeI, trailId oree.TrailId, stepId oree.StepId, startTime time.Time, endTime time.Time) {
	trail, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailId))
	}
	step, status := trail.StepWithId(stepId)
	if status == oree.NotFound {
		common.PrintLines(common.FormatIdNotFound("step", stepId))
	}
	o.Sessions().Create(oree.Session{
		StartTime: startTime,
		Duration:  endTime.Sub(startTime),
		Trail:     trail,
		Step:      step,
	})
	list(o)
}
