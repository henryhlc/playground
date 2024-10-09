package sessions

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  `update SessionId (StartTime| "_") (EndTime | "_") (TrailId | "_") (StepId | "_")`,
		Args: cobra.RangeArgs(5, 5),
		Run: func(cmd *cobra.Command, args []string) {
			runWithOree(func(o oree.OreeI) {
				update(o, args[0], args[1], args[2], args[3], args[4])
			})
		},
	}
	return cmd
}

func update(o oree.OreeI, sessionIdArg, startTimeArg, endTimeArg, trailIdArg, stepIdArg string) {
	session, ok := o.Sessions().WithId(oree.SessionId(sessionIdArg))
	if !ok {
		common.PrintLines(common.FormatIdNotFound("session", sessionIdArg))
		return
	}
	data, _ := session.Data()
	var trailId oree.TrailId
	if trailIdArg != "_" {
		trailId = oree.TrailId(trailIdArg)
	} else {
		trailId = data.Trail.Id()
	}
	trail, ok := o.Trails().WithId(trailId)
	if !ok {
		common.PrintLines(common.FormatIdNotFound("trail", trailIdArg))
		return
	}
	var stepId oree.StepId
	if stepIdArg != "_" {
		stepId = oree.StepId(stepIdArg)
	} else {
		stepId = data.Step.Id()
	}
	step, status := trail.StepWithId(stepId)
	if status == oree.NotFound {
		common.PrintLines(common.FormatIdNotFound("step", stepIdArg))
		return
	}

	var startTime time.Time
	var useArg bool
	if startTimeArg != "_" {
		startTime, useArg = common.TimeArg([]string{startTimeArg}, 0)
		if !useArg {
			fmt.Printf("Invalid StartTime %v\n", startTimeArg)
		}
	} else {
		startTime = data.StartTime
	}

	var endTime time.Time
	if endTimeArg != "_" {
		endTime, useArg = common.TimeArg([]string{endTimeArg}, 0)
		if !useArg {
			fmt.Printf("Invalid EndTime %v\n", endTimeArg)
		}
	} else {
		endTime = data.StartTime.Add(data.Duration)
	}

	session.Update(oree.Session{
		StartTime: startTime,
		Duration:  endTime.Sub(startTime),
		Trail:     trail,
		Step:      step,
	})

	list(o)
}
