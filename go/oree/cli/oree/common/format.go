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

func FormatNofM(n, m int, suffix string) []string {
	return []string{fmt.Sprintf("%v of %v %v", n, m, suffix)}
}

func PrintLines(lineLists ...[]string) {
	for _, lines := range lineLists {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
