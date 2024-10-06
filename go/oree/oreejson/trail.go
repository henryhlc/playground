package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailJD struct {
	Description   string   `json:"description"`
	ActiveSteps   *StepsJD `json:"activeSteps"`
	PinnedSteps   *StepsJD `json:"pinnedSteps"`
	ArchivedSteps *StepsJD `json:"archivedSteps"`
}

type TrailOJ struct {
	*TrailJD
	oreeJson OreeJson
	id       oree.TrailId
}

func NewTrailJD(trail oree.Trail) *TrailJD {
	return &TrailJD{
		Description:   trail.Description,
		ActiveSteps:   NewStepsJD(),
		ArchivedSteps: NewStepsJD(),
		PinnedSteps:   NewStepsJD(),
	}
}

func (t *TrailJD) EnsureInitialized() {
	if t.ActiveSteps == nil {
		t.ActiveSteps = NewStepsJD()
	}
	if t.ArchivedSteps == nil {
		t.ArchivedSteps = NewStepsJD()
	}
	if t.PinnedSteps == nil {
		t.PinnedSteps = NewStepsJD()
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

func (t TrailOJ) StepsWithStatus(status oree.StepStatus) oree.StepsI {
	switch status {
	case oree.Active:
		return t.ActiveSteps()
	case oree.Pinned:
		return t.PinnedSteps()
	case oree.Archived:
		return t.ArchivedSteps()
	}
	return StepsOJ{}
}

func (t TrailOJ) StepWithId(id oree.StepId) (oree.StepI, oree.StepStatus) {
	stepI, ok := t.ActiveSteps().WithId(id)
	if ok {
		return stepI, oree.Active
	}
	stepI, ok = t.PinnedSteps().WithId(id)
	if ok {
		return stepI, oree.Pinned
	}
	stepI, ok = t.ArchivedSteps().WithId(id)
	if ok {
		return stepI, oree.Archived
	}
	return StepOJ{}, oree.NotFound
}

func (t TrailOJ) ActiveSteps() oree.StepsI {
	return StepsFromData(t.TrailJD.ActiveSteps, t.oreeJson, t, oree.Active)
}

func (t TrailOJ) PinnedSteps() oree.StepsI {
	return StepsFromData(t.TrailJD.PinnedSteps, t.oreeJson, t, oree.Pinned)
}

func (t TrailOJ) ArchivedSteps() oree.StepsI {
	return StepsFromData(t.TrailJD.ArchivedSteps, t.oreeJson, t, oree.Archived)
}
