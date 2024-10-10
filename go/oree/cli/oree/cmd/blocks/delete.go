package blocks

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete BlockId...",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runWithOree(func(o oree.OreeI) {
				delete(o, args)
			})
		},
	}
	return cmd
}

func delete(o oree.OreeI, blockIds []string) {
	blocks := make([]oree.BlockI, len(blockIds))
	for i, blockId := range blockIds {
		block, ok := o.Blocks().WithId(oree.BlockId(blockId))
		if !ok {
			common.PrintLines(
				common.FormatIdNotFound("block", blockId),
			)
			return
		}
		blocks[i] = block
	}

	for _, block := range blocks {
		o.Blocks().Delete(block)
	}
	list(o)
}
