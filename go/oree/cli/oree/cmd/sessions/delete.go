package sessions

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete SessionId",
		Args: cobra.RangeArgs(1, 1),
		Run: func(cmd *cobra.Command, args []string) {
			sessionId, _ := common.StringArg(args, 0)

			runWithOree(func(o oree.OreeI) {
				delete(o, oree.SessionId(sessionId))
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, sessionId oree.SessionId) {
	session, ok := o.Sessions().WithId(sessionId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("session", sessionId),
		)
		return
	}
	o.Sessions().Delete(session)
	list(o)
}
