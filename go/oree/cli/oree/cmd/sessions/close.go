package sessions

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewCloseCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "close [EndTime]",
		Args: cobra.RangeArgs(1, 1),
		Run: func(cmd *cobra.Command, args []string) {
			endTime, useArg := common.TimeArg(args, 0)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[0])
			}
			runWithOree(func(o oree.OreeI) {
				close(o, endTime)
			})
		},
	}
	return cmd
}

func close(o oree.OreeI, endTime time.Time) {
	o.OpenSessionManager().Close(endTime)
	list(o)
}
