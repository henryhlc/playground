package areas

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update AreaId Description",
		Args: cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)
			description, useArg := common.StringArg(args, 1)
			if !useArg {
				fmt.Printf("Usage: %v", cmd.Use)
				return
			}

			runWithOree(func(o oree.OreeI) {
				update(o, oree.AreaId(areaId), description)
			})
		},
	}
	return cmd
}

func update(o oree.OreeI, areaId oree.AreaId, description string) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("area", areaId),
		)
		return
	}
	area.Update(oree.Area{
		Description: description,
	})
	common.PrintLines(
		common.FormatAreaWithTrails(area, defaultTrailsListLength),
	)
}
