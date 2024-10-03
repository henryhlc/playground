package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailJD struct {
	Description string   `json:"description"`
	ActiveSteps *StepsJD `json:"activeSteps"`
}

type TrailOJ struct {
	*TrailJD
	oreeJson OreeJson
	id       oree.TrailId
}

func NewTrailJD(trail oree.Trail) *TrailJD {
	return &TrailJD{
		Description: trail.Description,
		ActiveSteps: NewStepsJD(),
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

func (t TrailOJ) ActiveSteps() oree.StepsI {
	return StepsFromData(t.TrailJD.ActiveSteps, t.oreeJson)
}
