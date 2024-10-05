package oreejson

import (
	"strings"

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

func (oj *OreeJD) EnsureInitialized() {
	if oj.TrailsData == nil {
		oj.TrailsData = NewTrailsJD()
	}
}

func FromData(d *OreeJD) OreeJson {
	d.EnsureInitialized()
	return OreeJson{
		OreeJD: d,
	}
}

const digitMapping = "abcdefghijklmnopqrstuvwxyz012345789"

func (o OreeJson) getAndIncId() string {
	id := o.NextId
	o.NextId++
	var b strings.Builder
	if id == 0 {
		b.WriteByte(digitMapping[0])
		return b.String()
	}
	for id > 0 {
		b.WriteByte(digitMapping[id%36])
		id /= 36
	}
	return b.String()
}

func (o OreeJson) Trails() oree.TrailsI {
	return TrailsFromData(o.TrailsData, o)
}
