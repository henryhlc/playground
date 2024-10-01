package oree_test

import (
	"testing"

	"github.com/henryhlc/playground/go/oree"
	"github.com/henryhlc/playground/go/oree/oreejson"
)

func TestTrailsI(t *testing.T) {
	oreeI := oreejson.New()
	expectedDescription := "abc"
	id := oreeI.Trails().CreateTrail(oree.Trail{
		Description: expectedDescription,
	})
	actualDescription := oreeI.Trails().TrailWithId(id).Data().Description
	if actualDescription != expectedDescription {
		t.Errorf("Description expected: %v, actual: %v", expectedDescription, actualDescription)
	}
}
