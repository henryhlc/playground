package oreejson

import "github.com/henryhlc/playground/go/oree"

type StepJD struct {
	Description string `json:"description"`
}

type StepOJ struct {
	*StepJD
	id       oree.StepId
	oreeJson OreeJson
}

func NewStepJD(step oree.Step) *StepJD {
	return &StepJD{
		Description: step.Description,
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
