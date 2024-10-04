package trails

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete TrailId",
		Args: cobra.RangeArgs(1, 1),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)

			runWithOree(func(o oree.OreeI) {
				delete(o, oree.TrailId(trailId))
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, trailId oree.TrailId) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		fmt.Printf("No trail found for the given trail id \"%v\".\n", trailId)
		return
	}
	o.Trails().Delete(trailI)
	list(o)
}
