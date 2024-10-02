package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/oreejson"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "oree",
}

var jsonDataFileName string

func init() {
	rootCmd.PersistentFlags().StringVar(&jsonDataFileName, "json-data-file", "./oree.json", "path to JSON data file")
	rootCmd.AddCommand(trailsCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runFuncWithOree(f func(oree.OreeI)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(jsonDataFileName, os.O_RDWR|os.O_CREATE, 0644)
		maybePrintAndExit(err)
		defer file.Close()
		bytes, err := io.ReadAll(file)
		maybePrintAndExit(err)

		var oj oreejson.OreeJson

		var data oreejson.OreeJsonData
		if err := json.Unmarshal(bytes, &data); err != nil {
			oj = oreejson.FromData(oreejson.NewOreeJsonData())
		} else {
			oj = oreejson.FromData(&data)
		}

		f(oj)
		bytes, err = json.Marshal(oj.OreeJsonData)
		maybePrintAndExit(err)

		file.Truncate(0)
		file.WriteAt(bytes, 0)
	}
}

func maybePrintAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
