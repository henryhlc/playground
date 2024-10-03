package cmd

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/spf13/cobra"
)

var stepsCmd = &cobra.Command{
	Use: "steps",
	Run: runFuncWithOree(listSteps),
}

var stepTrailId string

var stepDescription string

var stepMoveId string
var stepMoveNextId string

var stepDeleteId string

var stepUpdateId string
var stepUpdateDescription string

func init() {
	stepsCmd.PersistentFlags().StringVar(&stepTrailId, "trail", "", "")
	stepsCmd.MarkFlagRequired("trail")

	stepsCmd.AddCommand(stepsPrependCmd)
	stepsCmd.AddCommand(stepsUpdateCmd)
	stepsCmd.AddCommand(stepsMoveCmd)
	stepsCmd.AddCommand(stepsDeleteCmd)

	stepsPrependCmd.Flags().StringVar(&stepDescription, "description", "", "Description of the trail.")
	stepsPrependCmd.MarkFlagRequired("description")

	stepsUpdateCmd.Flags().StringVar(&stepUpdateId, "id", "", "")
	stepsUpdateCmd.Flags().StringVar(&stepUpdateDescription, "description", "", "")
	stepsUpdateCmd.MarkFlagRequired("id")
	stepsUpdateCmd.MarkFlagRequired("description")

	stepsMoveCmd.Flags().StringVar(&stepMoveId, "id", "", "Id of the trail to move")
	stepsMoveCmd.Flags().StringVar(&stepMoveNextId, "next-id", "", "Id of the trail to move before to")

	stepsDeleteCmd.Flags().StringVar(&stepDeleteId, "id", "", "Id of the trail to be deleted")
}

func stepTrail(o oree.OreeI) (oree.TrailI, bool) {
	return o.Trails().WithId(oree.TrailId(stepTrailId))
}

func listSteps(o oree.OreeI) {
	trail, ok := stepTrail(o)
	if !ok {
		return
	}
	maxToList := 15
	steps := trail.ActiveSteps().FirstN(maxToList)
	for _, step := range steps {
		fmt.Printf("[%v] %v\n", step.Id(), step.Data().Description)
	}
	fmt.Println()
	fmt.Printf("%v out of %v steps\n", len(steps), trail.ActiveSteps().Len())
}

var stepsPrependCmd = &cobra.Command{
	Use: "prepend",
	Run: runFuncWithOree(prependStep),
}

func prependStep(o oree.OreeI) {
	trail, ok := stepTrail(o)
	if !ok {
		return
	}
	defer listSteps(o)
	trail.ActiveSteps().CreateFront(oree.Step{
		Description: stepDescription,
	})
}

var stepsMoveCmd = &cobra.Command{
	Use: "move",
	Run: runFuncWithOree(moveStep),
}

func moveStep(o oree.OreeI) {
	trail, ok := stepTrail(o)
	if !ok {
		return
	}
	defer listSteps(o)
	s, ok := trail.ActiveSteps().WithId(oree.StepId(stepMoveId))
	if !ok {
		return
	}
	nbr, ok := trail.ActiveSteps().WithId(oree.StepId(stepMoveNextId))
	if !ok {
		return
	}
	trail.ActiveSteps().PlaceBefore(s, nbr)
}

var stepsUpdateCmd = &cobra.Command{
	Use: "update",
	Run: runFuncWithOree(updateStep),
}

func updateStep(o oree.OreeI) {
	trail, ok := stepTrail(o)
	if !ok {
		return
	}
	defer listSteps(o)
	s, ok := trail.ActiveSteps().WithId(oree.StepId(stepUpdateId))
	if !ok {
		return
	}
	s.Update(oree.Step{
		Description: stepUpdateDescription,
	})
}

var stepsDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: runFuncWithOree(deleteStep),
}

func deleteStep(o oree.OreeI) {
	trail, ok := stepTrail(o)
	if !ok {
		return
	}
	defer listSteps(o)
	s, ok := trail.ActiveSteps().WithId(oree.StepId(stepDeleteId))
	if !ok {
		return
	}
	trail.ActiveSteps().Delete(s)
}
