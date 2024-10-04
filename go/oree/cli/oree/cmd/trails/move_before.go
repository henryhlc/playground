package trails

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewMoveBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "move-before TrailId NeighborTrailId",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			trailId, _ := common.StringArg(args, 0)
			nbrTrailId, _ := common.StringArg(args, 1)

			runWithOree(func(o oree.OreeI) {
				moveBefore(o, oree.TrailId(trailId), oree.TrailId(nbrTrailId))
			})
		},
	}
	return cmd
}

func moveBefore(o oree.OreeI, trailId oree.TrailId, nbrTrailId oree.TrailId) {
	trailI, ok := o.Trails().WithId(trailId)
	if !ok {
		fmt.Printf("No trail found for the given trail id \"%v\".\n", trailId)
		return
	}
	nbrTrailI, ok := o.Trails().WithId(nbrTrailId)
	if !ok {
		fmt.Printf("No trail found for the given neighbor trail id \"%v\".\n", nbrTrailId)
		return
	}
	o.Trails().PlaceBefore(trailI, nbrTrailI)
	list(o)
}
