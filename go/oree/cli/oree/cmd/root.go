package cmd

import (
	"time"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/areas"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/blocks"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/sessions"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/steps"
	"github.com/henryhlc/playground/go/oree/cli/oree/cmd/trails"
	"github.com/henryhlc/playground/go/oree/cli/oree/common"
	"github.com/spf13/cobra"
)

const jsonDataFileFlag = "json-data-file"
const defaultDashTrailsN = 10
const defaultDashPinnedStepsN = 3
const defaultDashActiveStepsN = 5
const defaultDashAreaTrailsN = 30
const defaultDashBlocksN = 3

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
	cmd.AddCommand(sessions.NewCmd(runWithOree))
	cmd.AddCommand(blocks.NewCmd(runWithOree))

	return cmd
}

func dash(o oree.OreeI) {
	trails := o.Trails().FirstN(defaultDashTrailsN)
	lines := []string{}

	if os, ok := o.OpenSessionManager().Data(); ok {
		lines = common.ConcatLines(lines,
			common.FormatOpenSession(os))
	}

	currentBlock, blockExists := o.Blocks().LastBlockCovering(time.Now())
	if blockExists {
		titleLines := []string{}
		if len(lines) > 0 {
			titleLines = append(titleLines, "")
		}
		titleLines = append(titleLines, "Current block")
		lines = common.ConcatLines(lines,
			titleLines,
			common.FormatBlock(currentBlock),
		)
		blockData, ok := currentBlock.Data()
		if ok {
			switch at := blockData.Context.(type) {
			case oree.AreaI:
				trails := at.Trails().FirstN(30)
				for _, trail := range trails {
					lines = common.ConcatLines(lines, common.FormatPrefix("  ", common.ConcatLines(
						common.FormatTrail(trail),
						common.FormatPrefix("  ", common.FormatSteps(trail.StepsWithStatus(oree.Pinned).FirstN(defaultDashPinnedStepsN))),
						common.FormatPrefix("  ", common.FormatSteps(trail.StepsWithStatus(oree.Active).FirstN(defaultDashActiveStepsN))),
					)))
				}
			case oree.TrailI:
				lines = common.ConcatLines(
					lines,
					common.FormatTrail(at),
					common.FormatPrefix("  ", common.FormatSteps(at.StepsWithStatus(oree.Pinned).FirstN(defaultDashPinnedStepsN))),
					common.FormatPrefix("  ", common.FormatSteps(at.StepsWithStatus(oree.Active).FirstN(defaultDashActiveStepsN))),
				)
			}
		}
	} else {
		var blocks []oree.BlockI
		blockBefore, ok := o.Blocks().LastBlockStartBefore(time.Now())
		if !ok {
			blocks = o.Blocks().LastN(defaultDashBlocksN)
		} else {
			blocks = o.Blocks().NAfter(defaultDashBlocksN, blockBefore)
		}
		var titleLines []string
		if len(lines) > 0 {
			titleLines = append(titleLines, "")
		}
		titleLines = append(titleLines, "Upcoming blocks")
		if len(blocks) > 0 {
			lines = common.ConcatLines(lines,
				titleLines,
				common.FormatBlocks(blocks),
			)
		}

		titleLines = []string{}
		if len(lines) > 0 {
			titleLines = append(titleLines, "")
		}
		titleLines = append(titleLines, "Trails")
		lines = common.ConcatLines(lines, titleLines)
		for _, trail := range trails {
			lines = common.ConcatLines(
				lines,
				common.FormatTrail(trail),
				common.FormatPrefix("  ", common.FormatSteps(trail.StepsWithStatus(oree.Pinned).FirstN(defaultDashPinnedStepsN))),
				common.FormatPrefix("  ", common.FormatSteps(trail.StepsWithStatus(oree.Active).FirstN(defaultDashActiveStepsN))),
			)
		}
	}

	common.PrintLines(lines)
}
