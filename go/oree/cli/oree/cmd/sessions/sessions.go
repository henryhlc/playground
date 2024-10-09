package sessions

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListLength = 15

func NewCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "sessions [n]",
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			n, _ := common.IntArgOrDefault(args, 0, defaultListLength)
			runWithOree(func(o oree.OreeI) {
				listN(o, n)
			})
		},
	}
	cmd.AddCommand(NewCreateCmd(runWithOree))
	cmd.AddCommand(NewDeleteCmd(runWithOree))
	cmd.AddCommand(NewListBeforeCmd(runWithOree))
	cmd.AddCommand(NewUpdateCmd(runWithOree))
	return cmd
}

func list(o oree.OreeI) {
	listN(o, defaultListLength)
}

func listN(o oree.OreeI, n int) {
	sessions := o.Sessions().LastN(n)
	common.PrintLines(
		common.FormatSessions(sessions),
		[]string{""},
		common.FormatNofM(len(sessions), o.Sessions().Len(), "sessions"),
	)
}
