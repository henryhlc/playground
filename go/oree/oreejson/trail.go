package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailJsonData struct {
	Description string `json:"description"`
}

type TrailJson struct {
	*TrailJsonData
	oreeJson OreeJson
	id       oree.TrailId
}

func NewTrailJsonData(trail oree.Trail) *TrailJsonData {
	return &TrailJsonData{
		Description: trail.Description,
	}
}

func (t TrailJson) Id() oree.TrailId {
	return t.id
}

func (t TrailJson) Data() oree.Trail {
	return oree.Trail{
		Description: t.Description,
	}
}

func (t TrailJson) Update(trail oree.Trail) {
	t.Description = trail.Description
}
