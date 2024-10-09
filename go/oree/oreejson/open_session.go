package oreejson

import (
	"time"

	"github.com/henryhlc/playground/go/oree"
)

type OpenSessionJD struct {
	StartTime time.Time    `json:"startTime"`
	StepId    oree.StepId  `json:"stepId"`
	TrailId   oree.TrailId `json:"trailId"`
}

type OpenSessionManagerJD struct {
	OpenSession *OpenSessionJD `json:"session"`
}

type OpenSessionManagerOJ struct {
	*OpenSessionManagerJD
	oreeJson OreeJson
}

func NewOpenSessionManagerJD() *OpenSessionManagerJD {
	return &OpenSessionManagerJD{}
}

func OpenSessionManagerFromData(d *OpenSessionManagerJD, oj OreeJson) OpenSessionManagerOJ {
	return OpenSessionManagerOJ{
		OpenSessionManagerJD: d,
		oreeJson:             oj,
	}
}

func (os OpenSessionManagerOJ) Data() (oree.OpenSession, bool) {
	if os.OpenSession == nil {
		return oree.OpenSession{}, false
	}
	trail, ok := os.oreeJson.Trails().WithId(os.OpenSession.TrailId)
	if !ok {
		return oree.OpenSession{}, false
	}
	step, status := trail.StepWithId(os.OpenSession.StepId)
	if status == oree.NotFound {
		return oree.OpenSession{}, false
	}
	return oree.OpenSession{
		StartTime: os.OpenSession.StartTime,
		Trail:     trail,
		Step:      step,
	}, true
}

func (os OpenSessionManagerOJ) Open(session oree.OpenSession) {
	os.OpenSession = &OpenSessionJD{
		StartTime: session.StartTime,
		TrailId:   session.Trail.Id(),
		StepId:    session.Step.Id(),
	}
}

func (os OpenSessionManagerOJ) Close(endTime time.Time) {
	data, ok := os.Data()
	if !ok {
		os.OpenSession = nil
		return
	}
	os.oreeJson.Sessions().Create(oree.Session{
		StartTime: data.StartTime,
		Duration:  endTime.Sub(data.StartTime),
		Trail:     data.Trail,
		Step:      data.Step,
	})
	os.OpenSession = nil
}
