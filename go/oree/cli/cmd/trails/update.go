package trails

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/common"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update TrailId Description",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			description, useArg := common.StringArg(args, 1)
			if !useArg {
				fmt.Printf("Usage: %v", cmd.Use)
				return
			}

			runWithOree(func(o oree.OreeI) {
				update(o, oree.TrailId(trailId), description)
			})
		},
	}
	return cmd
}

func update(o oree.OreeI, trailId oree.TrailId, description string) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		fmt.Printf("No trail found for the given trail id \"%v\".\n", trailId)
		return
	}
	trailI.Update(oree.Trail{
		Description: description,
	})
	common.PrintLines(
		common.FormatTrail(trailI),
	)
}
