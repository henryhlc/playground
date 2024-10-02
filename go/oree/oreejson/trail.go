package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailJD struct {
	Description string `json:"description"`
}

type TrailOJ struct {
	*TrailJD
	oreeJson OreeJson
	id       oree.TrailId
}

func NewTrailJD(trail oree.Trail) *TrailJD {
	return &TrailJD{
		Description: trail.Description,
	}
}

func (t TrailOJ) Id() oree.TrailId {
	return t.id
}

func (t TrailOJ) Data() oree.Trail {
	return oree.Trail{
		Description: t.Description,
	}
}

func (t TrailOJ) Update(trail oree.Trail) {
	t.Description = trail.Description
}
