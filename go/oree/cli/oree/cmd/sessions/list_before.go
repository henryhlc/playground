package sessions

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListBeforeLength = 15

func NewListBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-before SessionId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			sessionId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListBeforeLength)

			runWithOree(func(o oree.OreeI) {
				listNBefore(o, oree.SessionId(sessionId), n)
			})
		},
	}
	return cmd
}

func listNBefore(o oree.OreeI, sessionId oree.SessionId, n int) {
	session, ok := o.Sessions().WithId(sessionId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("session", sessionId),
		)
		return
	}
	sessions := o.Sessions().NBefore(n, session)
	numSessions := len(sessions)
	total := o.Sessions().Len()
	common.PrintLines(
		common.FormatSessions(sessions),
		[]string{""},
		common.FormatNofM(numSessions, total, "sessions"),
	)
}
