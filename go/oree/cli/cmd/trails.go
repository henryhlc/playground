package cmd

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
	"github.com/spf13/cobra"
)

var trailsCmd = &cobra.Command{
	Use: "trails",
	Run: runFuncWithOree(listTrails),
}

var trailDescription string

var trailMoveId string
var trailMoveNextId string

var trailDeleteId string

var trailUpdateId string
var trailUpdateDescription string

func init() {
	trailsCmd.AddCommand(trailsPrependCmd)
	trailsCmd.AddCommand(trailsUpdateCmd)
	trailsCmd.AddCommand(trailsMoveCmd)
	trailsCmd.AddCommand(trailsDeleteCmd)

	trailsPrependCmd.Flags().StringVar(&trailDescription, "description", "", "Description of the trail.")
	trailsPrependCmd.MarkFlagRequired("description")

	trailsUpdateCmd.Flags().StringVar(&trailUpdateId, "id", "", "")
	trailsUpdateCmd.Flags().StringVar(&trailUpdateDescription, "description", "", "")
	trailsUpdateCmd.MarkFlagRequired("id")
	trailsUpdateCmd.MarkFlagRequired("description")

	trailsMoveCmd.Flags().StringVar(&trailMoveId, "id", "", "Id of the trail to move")
	trailsMoveCmd.Flags().StringVar(&trailMoveNextId, "next-id", "", "Id of the trail to move before to")

	trailsDeleteCmd.Flags().StringVar(&trailDeleteId, "id", "", "Id of the trail to be deleted")
}

func listTrails(o oree.OreeI) {
	maxToList := 15
	trails := o.Trails().FirstN(maxToList)
	for _, trail := range trails {
		fmt.Printf("[%v] %v\n", trail.Id(), trail.Data().Description)
	}
	fmt.Println()
	fmt.Printf("%v out of %v trails\n", len(trails), o.Trails().Len())
}

var trailsPrependCmd = &cobra.Command{
	Use: "prepend",
	Run: runFuncWithOree(prependTrail),
}

func prependTrail(o oree.OreeI) {
	defer listTrails(o)
	o.Trails().CreateFront(oree.Trail{
		Description: trailDescription,
	})
}

var trailsMoveCmd = &cobra.Command{
	Use: "move",
	Run: runFuncWithOree(moveTrail),
}

func moveTrail(o oree.OreeI) {
	defer listTrails(o)
	t, ok := o.Trails().WithId(oree.TrailId(trailMoveId))
	if !ok {
		return
	}
	nbr, ok := o.Trails().WithId(oree.TrailId(trailMoveNextId))
	if !ok {
		return
	}
	o.Trails().PlaceBefore(t, nbr)
}

var trailsUpdateCmd = &cobra.Command{
	Use: "update",
	Run: runFuncWithOree(updateTrail),
}

func updateTrail(o oree.OreeI) {
	defer listTrails(o)
	t, ok := o.Trails().WithId(oree.TrailId(trailUpdateId))
	if !ok {
		return
	}
	o.Trails().Update(t, oree.Trail{
		Description: trailUpdateDescription,
	})
}

var trailsDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: runFuncWithOree(deleteTrail),
}

func deleteTrail(o oree.OreeI) {
	defer listTrails(o)
	t, ok := o.Trails().WithId(oree.TrailId(trailDeleteId))
	if !ok {
		return
	}
	o.Trails().Delete(t)
}
