package oreejson

import (
	"time"

	"github.com/henryhlc/playground/go/oree"
)

type SessionJD struct {
	StartTime time.Time     `json:"startTime"`
	Duration  time.Duration `json:"duration"`
	TrailId   oree.TrailId  `json:"trailId"`
	StepId    oree.StepId   `json:"stepId"`
}

type SessionOJ struct {
	*SessionJD
	oreeJson OreeJson
	id       oree.SessionId
}

func NewSessionJD(s oree.Session) *SessionJD {
	return &SessionJD{
		StartTime: s.StartTime,
		Duration:  s.Duration,
		TrailId:   s.Trail.Id(),
		StepId:    s.Step.Id(),
	}
}

func SessionFromData(d *SessionJD, oj OreeJson, id oree.SessionId) SessionOJ {
	return SessionOJ{
		SessionJD: d,
		oreeJson:  oj,
		id:        id,
	}
}

func (s SessionOJ) Id() oree.SessionId {
	return s.id
}

func (s SessionOJ) Data() (oree.Session, bool) {
	t, ok := s.oreeJson.Trails().WithId(s.TrailId)
	if !ok {
		return oree.Session{}, false
	}
	st, status := t.StepWithId(s.StepId)
	if status == oree.NotFound {
		return oree.Session{}, false
	}

	return oree.Session{
		StartTime: s.StartTime,
		Duration:  s.Duration,
		Step:      st,
		Trail:     t,
	}, true
}

func (s SessionOJ) Update(session oree.Session) {
	s.StartTime = session.StartTime
	s.Duration = session.Duration
	s.TrailId = session.Trail.Id()
	s.StepId = session.Step.Id()
	sessions := s.oreeJson.Sessions().(SessionsOJ)
	sessions.FixOrder(ListItem[oree.SessionId, SessionJD]{
		Id:   s.id,
		Elem: s.SessionJD,
	})
}
