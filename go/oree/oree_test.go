package oree_test

import (
	"fmt"
	"testing"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/oreejson"
)

func TestOreeI(t *testing.T) {
	oreeI := oreejson.New()
	id := oreeI.CreateTrail(oree.Trail{
		Description: "abc",
	})
	fmt.Println(id)
}

func TestOree(t *testing.T) {
	o := oree.Oree{
		OreeI: oreejson.New(),
	}
	id := o.CreateTrail(oree.Trail{
		Description: "abc",
	})
	fmt.Println(id)

}
