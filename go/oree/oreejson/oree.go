package oreejson

import (
	"strconv"

	"github.com/henryhlc/playground/go/oree"
)

type OreeJsonData struct {
	NextId     int             `json:"nextId"`
	TrailsData *TrailsJsonData `json:"trails"`
}

func NewOreeJsonData() *OreeJsonData {
	return &OreeJsonData{
		NextId:     0,
		TrailsData: NewTrailsJsonData(),
	}
}

type OreeJson struct {
	*OreeJsonData
}

func New() OreeJson {
	return OreeJson{
		OreeJsonData: NewOreeJsonData(),
	}
}

func NewFromOreeJsonData(d *OreeJsonData) OreeJson {
	return OreeJson{
		OreeJsonData: d,
	}
}

func (o OreeJson) getAndIncId() string {
	id := o.NextId
	o.NextId++
	return strconv.Itoa(id)
}

func (o OreeJson) Trails() oree.TrailsI {
	return TrailsJson{
		TrailsJsonData: o.TrailsData,
		oreeJson:       o,
	}
}
