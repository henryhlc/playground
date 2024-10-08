package areas

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete AreaId",
		Args: cobra.RangeArgs(1, 1),
		Run: func(cmd *cobra.Command, args []string) {
			areaId, _ := common.StringArg(args, 0)

			runWithOree(func(o oree.OreeI) {
				delete(o, oree.AreaId(areaId))
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, areaId oree.AreaId) {
	area, ok := o.Areas().WithId(areaId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("area", areaId),
		)
		return
	}
	o.Areas().Delete(area)
	list(o)
}
