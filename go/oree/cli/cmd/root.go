package cmd

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/cmd/trails"
	"github.com/henryhlc/playground/go/oree/cli/common"
	"github.com/spf13/cobra"
)

const jsonDataFileFlag = "json-data-file"

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "oree",
	}

	jsonDataFilePath := cmd.PersistentFlags().String(jsonDataFileFlag, "./oree.json", "Path to JSON data file")
	cmd.MarkFlagFilename(jsonDataFileFlag)

	runWithOree := func(f func(oree.OreeI)) {
		common.RunWithOree(*jsonDataFilePath, f)
	}
	cmd.AddCommand(trails.NewCmd(runWithOree))

	return cmd
}
