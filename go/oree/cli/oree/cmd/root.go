package cmd

import (
	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/areas"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/steps"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/trails"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const jsonDataFileFlag = "json-data-file"
const defaultDashTrailsN = 10
const defaultDashStepsN = 3

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "oree",
	}

	jsonDataFilePath := cmd.PersistentFlags().String(jsonDataFileFlag, "./oree.json", "Path to JSON data file")
	cmd.MarkFlagFilename(jsonDataFileFlag)

	runWithOree := func(f func(oree.OreeI)) {
		common.RunWithOree(*jsonDataFilePath, f)
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		runWithOree(dash)
	}
	cmd.AddCommand(trails.NewCmd(runWithOree))
	cmd.AddCommand(steps.NewCmd(runWithOree))
	cmd.AddCommand(areas.NewCmd(runWithOree))

	return cmd
}

func dash(o oree.OreeI) {
	trails := o.Trails().FirstN(defaultDashTrailsN)
	lines := []string{}
	for _, trail := range trails {
		lines = common.ConcatLines(
			lines,
			common.FormatTrail(trail),
			common.FormatPrefix("  ", common.FormatSteps(trail.StepsWithStatus(oree.Active).FirstN(defaultDashStepsN))),
		)
	}
	lines = common.ConcatLines(
		lines,
		[]string{""},
		common.FormatNofM(len(trails), o.Trails().Len(), "trails"))
	common.PrintLines(lines)
}
