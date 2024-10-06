package oreejson

import "github.com/henryhlc/playground/go/oree"

type StepJD struct {
	Description string `json:"description"`
}

type StepOJ struct {
	*StepJD
	id       oree.StepId
	status   oree.StepStatus
	trail    TrailOJ
	oreeJson OreeJson
}

func NewStepJD(step oree.Step) *StepJD {
	return &StepJD{
		Description: step.Description,
	}
}

func StepFromData(d *StepJD, oj OreeJson, t TrailOJ, id oree.StepId, s oree.StepStatus) StepOJ {
	return StepOJ{
		StepJD:   d,
		id:       id,
		status:   s,
		trail:    t,
		oreeJson: oj,
	}

}

func (s StepOJ) Id() oree.StepId {
	return s.id
}

func (s StepOJ) Data() oree.Step {
	return oree.Step{
		Description: s.Description,
	}
}

func (s StepOJ) Update(step oree.Step) {
	s.Description = step.Description
}

func (s StepOJ) UpdateStatus(targetStatus oree.StepStatus) {
	if s.status == targetStatus {
		return
	}
	s.trail.StepsWithStatus(s.Status()).Delete(s)
	s.trail.StepsWithStatus(targetStatus).(StepsOJ).PlaceItemFront(
		ListItem[oree.StepId, StepJD]{
			Id:   s.Id(),
			Elem: s.StepJD,
		},
	)
}

func (s StepOJ) Status() oree.StepStatus {
	return s.status
}
