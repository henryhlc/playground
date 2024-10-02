package oreejson

import (
	"strconv"

	"github.com/henryhlc/playground/go/oree"
)

type OreeJD struct {
	NextId     int       `json:"nextId"`
	TrailsData *TrailsJD `json:"trails"`
}

type OreeJson struct {
	*OreeJD
}

func NewOreeJD() *OreeJD {
	return &OreeJD{
		NextId:     0,
		TrailsData: NewTrailsJD(),
	}
}

func FromData(d *OreeJD) OreeJson {
	return OreeJson{
		OreeJD: d,
	}
}

func (o OreeJson) getAndIncId() string {
	id := o.NextId
	o.NextId++
	return strconv.Itoa(id)
}

func (o OreeJson) Trails() oree.TrailsI {
	return TrailsFromData(o.TrailsData, o)
}
