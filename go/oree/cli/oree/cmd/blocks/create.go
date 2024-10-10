package blocks

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewCreateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create StartTime EndTime TargetDuration Descripttion [(TrailId | AreaId)]",
		Args: cobra.RangeArgs(4, 5),
		Run: func(cmd *cobra.Command, args []string) {
			startTime, useArg := common.TimeArg(args, 0)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[0])
				return
			}
			endTime, useArg := common.TimeArg(args, 1)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[1])
				return
			}
			targetDuration, useArg := common.DurationArg(args, 2)
			if !useArg {
				fmt.Printf("Invadid duration %v\n", args[2])
				return
			}
			description, _ := common.StringArg(args, 3)
			contextId, _ := common.StringArg(args, 4)
			runWithOree(func(o oree.OreeI) {
				create(o, startTime, endTime, targetDuration, description, contextId)
			})
		},
	}
	return cmd
}

func create(o oree.OreeI, startTime time.Time, endTime time.Time, targetDuration time.Duration, description string, contextId string) {
	var context interface{}
	if contextId != "" {
		trail, ok := o.Trails().WithId(oree.TrailId(contextId))
		if ok {
			context = trail
		}
		area, ok := o.Areas().WithId(oree.AreaId(contextId))
		if ok {
			context = area
		}
	}
	o.Blocks().Create(oree.Block{
		Description:    description,
		StartTime:      startTime,
		Duration:       endTime.Sub(startTime),
		TargetDuration: targetDuration,
		Context:        context,
	})
	list(o)
}
