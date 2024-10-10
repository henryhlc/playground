package blocks

import (
	"fmt"
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDupWithDateCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "dup-with-date Date BlockIds...",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			date, useArg := common.TimeArg(args, 0)
			if !useArg {
				fmt.Printf("Invalid time %v\n", args[0])
				return
			}
			runWithOree(func(o oree.OreeI) {
				dupWithDate(o, date, args[1:])
			})

		},
	}
	return cmd
}

func dupWithDate(o oree.OreeI, date time.Time, blockIds []string) {
	blocks := make([]oree.BlockI, len(blockIds))
	for i, id := range blockIds {
		block, ok := o.Blocks().WithId(oree.BlockId(id))
		if !ok {
			common.PrintLines(common.FormatIdNotFound("block", id))
			return
		}
		_, ok = block.Data()
		if !ok {
			fmt.Printf("Invalid block %v\n", id)
		}
		blocks[i] = block
	}
	for _, block := range blocks {
		data, _ := block.Data()

		data.StartTime = time.Date(date.Year(), date.Month(), date.Day(),
			data.StartTime.Hour(), data.StartTime.Minute(), 0, 0, time.Local)
		o.Blocks().Create(data)
	}
	list(o)
}
