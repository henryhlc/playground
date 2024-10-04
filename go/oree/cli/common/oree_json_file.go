package common

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/oreejson"
	"github.com/spf13/cobra"
)

func RunFuncWithOree(jsonFilePath string, f func(oree.OreeI)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		RunWithOree(jsonFilePath, f)
	}
}

func RunWithOree(jsonFilePath string, f func(oree.OreeI)) {
	file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
	maybePrintAndExit(err)
	defer file.Close()
	bytes, err := io.ReadAll(file)
	maybePrintAndExit(err)

	var oj oreejson.OreeJson

	var data oreejson.OreeJD
	if err := json.Unmarshal(bytes, &data); err != nil {
		oj = oreejson.FromData(oreejson.NewOreeJD())
	} else {
		oj = oreejson.FromData(&data)
	}

	f(oj)
	bytes, err = json.Marshal(oj.OreeJD)
	maybePrintAndExit(err)

	file.Truncate(0)
	file.WriteAt(bytes, 0)

}

func maybePrintAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
