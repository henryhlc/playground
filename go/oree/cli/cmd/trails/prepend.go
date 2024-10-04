package trails

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/common"
	"github.com/spf13/cobra"
)

func NewPrependCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "prepend Description",
		Args: cobra.RangeArgs(1, 1),
		Run: func(cmd *cobra.Command, args []string) {
			description, _ := common.StringArg(args, 0)
			runWithOree(func(o oree.OreeI) {
				prepend(o, description)
			})
		},
	}
	return cmd
}

func prepend(o oree.OreeI, description string) {
	o.Trails().CreateFront(oree.Trail{
		Description: description,
	})
	list(o)
}
