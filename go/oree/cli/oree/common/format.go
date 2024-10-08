package common

import (
	"fmt"

	"github.com/henryhlc/playground/go/oree"
)

func FormatTrail(t oree.TrailI) []string {
	return []string{
		fmt.Sprintf("[%v] %v", t.Id(), t.Data().Description),
	}
}

func FormatTrails(ts []oree.TrailI) []string {
	lines := []string{}
	for _, t := range ts {
		lines = append(lines, FormatTrail(t)...)
	}
	return lines
}

func FormatStep(s oree.StepI) []string {
	return []string{
		fmt.Sprintf("[%v] %v", s.Id(), s.Data().Description),
	}
}

func FormatSteps(ss []oree.StepI) []string {
	lines := []string{}
	for _, s := range ss {
		lines = append(lines, FormatStep(s)...)
	}
	return lines
}

func FormatStepsSection(status oree.StepStatus, total int, ss []oree.StepI) []string {
	var title string
	switch status {
	case oree.Active:
		title = "Active steps"
	case oree.Archived:
		title = "Archived steps"
	case oree.Pinned:
		title = "Pinned steps"
	}

	lines := []string{
		fmt.Sprintf("%v (%v of %v steps)", title, len(ss), total),
	}
	lines = append(lines, FormatPrefix("  ", FormatSteps(ss))...)
	return FormatPrefix("  ", lines)
}

func FormatTrailWithSteps(trail oree.TrailI, statuses []oree.StepStatus, n int) []string {
	lines := FormatTrail(trail)
	for _, status := range statuses {
		steps := trail.StepsWithStatus(status)
		lines = ConcatLines(lines, FormatStepsSection(status, steps.Len(), steps.FirstN(n)))
	}
	return lines
}

func FormatArea(area oree.AreaI) []string {
	return []string{fmt.Sprintf("[%v] %v", area.Id(), area.Data().Description)}
}

func FormatAreas(areas []oree.AreaI) []string {
	lines := []string{}
	for _, area := range areas {
		lines = ConcatLines(lines, FormatArea(area))
	}
	return lines
}

func FormatAreaWithTrails(area oree.AreaI, n int) []string {
	return ConcatLines(
		FormatArea(area),
		FormatPrefix("  ", FormatTrails(area.Trails().FirstN(n))),
	)
}

func FormatNofM(n, m int, suffix string) []string {
	return []string{fmt.Sprintf("%v of %v %v", n, m, suffix)}
}

func FormatIdNotFound(itemType string, id interface{}) []string {
	return []string{
		fmt.Sprintf("No %v found for the given id \"%v\".", itemType, id),
	}
}

func FormatPrefix(prefix string, lines []string) []string {
	linesWithPrefix := make([]string, len(lines))
	for i, line := range lines {
		linesWithPrefix[i] = prefix + line
	}
	return linesWithPrefix
}

func ConcatLines(lineLists ...[]string) []string {
	concatLines := []string{}
	for _, lines := range lineLists {
		concatLines = append(concatLines, lines...)
	}
	return concatLines
}

func PrintLines(lineLists ...[]string) {
	for _, lines := range lineLists {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
