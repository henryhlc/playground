package areas

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/areas/trails"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 15

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use: "areas",
		Run: func(cmd *cobra.Command, args []string) {
			runWithOree(list)
		},
	}
	cmd.AddCommand(trails.NewCmd(runWithOree))
	cmd.AddCommand(NewListAfterCmd(runWithOree))
	cmd.AddCommand(NewPrependCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	cmd.AddCommand(NewUpdateCmd(runWithOree))
	cmd.AddCommand(NewMoveBeforeCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI) {
	areas := o.Areas().FirstN(defaultListLength)
	common.PrintLines(
		common.FormatAreas(areas),
		[]string{""},
		common.FormatNofM(len(areas), o.Areas().Len(), "areas"),
	)
}
