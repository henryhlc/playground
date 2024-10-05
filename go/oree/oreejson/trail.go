package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailJD struct {
	Description    string   `json:"description"`
	ActiveSteps    *StepsJD `json:"activeSteps"`
	ArchivedSteps  *StepsJD `json:"archivedSteps"`
	HighlightSteps *StepsJD `json:"highlightSteps"`
}

type TrailOJ struct {
	*TrailJD
	oreeJson OreeJson
	id       oree.TrailId
}

func NewTrailJD(trail oree.Trail) *TrailJD {
	return &TrailJD{
		Description:    trail.Description,
		ActiveSteps:    NewStepsJD(),
		ArchivedSteps:  NewStepsJD(),
		HighlightSteps: NewStepsJD(),
	}
}

func (t *TrailJD) EnsureInitialized() {
	if t.ActiveSteps == nil {
		t.ActiveSteps = NewStepsJD()
	}
	if t.ArchivedSteps == nil {
		t.ArchivedSteps = NewStepsJD()
	}
	if t.HighlightSteps == nil {
		t.HighlightSteps = NewStepsJD()
	}
}

func TrailFromData(data *TrailJD, oj OreeJson, id oree.TrailId) TrailOJ {
	data.EnsureInitialized()
	return TrailOJ{
		TrailJD:  data,
		oreeJson: oj,
		id:       id,
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
