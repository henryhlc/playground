package blocks

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  `update BlockId (StartTime| "_") (EndTime | "_") (TargetDuration | "_") (Description | "_") [(TrailId | AreaId)]`,
		Args: cobra.RangeArgs(5, 6),
		Run: func(cmd *cobra.Command, args []string) {
			contextId, _ := common.StringArg(args, 5)
			runWithOree(func(o oree.OreeI) {
				update(o, args[0], args[1], args[2], args[3], args[4], contextId)
			})
		},
	}
	return cmd
}

func update(o oree.OreeI, blockIdArg, startTimeArg, endTimeArg, targetDurationArg, descriptionArg, contextIdArg string) {
	block, ok := o.Blocks().WithId(oree.BlockId(blockIdArg))
	if !ok {
		common.PrintLines(common.FormatIdNotFound("blockId", blockIdArg))
		return
	}
	data, ok := block.Data()
	if !ok {
		fmt.Printf("Invalid block %v\n", blockIdArg)
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

	var targetDuration time.Duration
	if targetDurationArg != "_" {
		targetDuration, useArg = common.DurationArg([]string{targetDurationArg}, 0)
		if !useArg {
			fmt.Printf("Invalid duration %v\n", targetDurationArg)
		}
	} else {
		targetDuration = data.TargetDuration
	}

	var description string
	if descriptionArg != "_" {
		description = descriptionArg
	} else {
		description = data.Description
	}

	var context interface{}
	switch contextIdArg {
	case "_":
		context = data.Context
	default:
		found := false
		area, ok := o.Areas().WithId(oree.AreaId(contextIdArg))
		if ok {
			found = true
			context = area
		}
		trail, ok := o.Trails().WithId(oree.TrailId(contextIdArg))
		if ok {
			found = true
			context = trail
		}
		if !found {
			common.PrintLines(common.FormatIdNotFound("trail or area", contextIdArg))
			return
		}
	}

	block.Update(oree.Block{
		Description:    description,
		StartTime:      startTime,
		Duration:       endTime.Sub(startTime),
		TargetDuration: targetDuration,
		Context:        context,
	})

	list(o)
}
