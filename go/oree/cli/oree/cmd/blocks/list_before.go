package blocks

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const defaultListBeforeLength = 15

func NewListBeforeCmd(runWithOree func(func(oree.OreeI))) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list-before BlockId [n]",
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			blockId, _ := common.StringArg(args, 0)
			n, _ := common.IntArgOrDefault(args, 1, defaultListBeforeLength)

			runWithOree(func(o oree.OreeI) {
				listNBefore(o, oree.BlockId(blockId), n)
			})
		},
	}
	return cmd
}

func listNBefore(o oree.OreeI, blockId oree.BlockId, n int) {
	block, ok := o.Blocks().WithId(blockId)
	if !ok {
		common.PrintLines(
			common.FormatIdNotFound("block", blockId),
		)
		return
	}
	blocks := o.Blocks().NBefore(n, block)
	numBlocks := len(blocks)
	total := o.Blocks().Len()
	common.PrintLines(
		common.FormatBlocks(blocks),
		[]string{""},
		common.FormatNofM(numBlocks, total, "blocks"),
	)
}
