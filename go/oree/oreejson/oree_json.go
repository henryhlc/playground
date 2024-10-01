package oreejson

import (
	"github.com/henryhlc/playground/go/oree"
)

type OreeJson struct {
}

func New() *OreeJson {
	return &OreeJson{}
}

func (o *OreeJson) CreateTrail(trail oree.Trail) oree.TrailId {
	return "123"
}
