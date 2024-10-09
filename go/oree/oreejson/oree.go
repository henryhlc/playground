package oreejson

import (
	"strings"

	"github.com/henryhlc/playground/go/oree"
)

type OreeJD struct {
	NextId       int         `json:"nextId"`
	TrailsData   *TrailsJD   `json:"trails"`
	AreasData    *AreasJD    `json:"areas"`
	SessionsData *SessionsJD `json:"sessions"`
}

type OreeJson struct {
	*OreeJD
}

func NewOreeJD() *OreeJD {
	return &OreeJD{
		NextId:       0,
		TrailsData:   NewTrailsJD(),
		AreasData:    NewAreasJD(),
		SessionsData: NewSessionsJD(),
	}
}

func (oj *OreeJD) EnsureInitialized() {
	if oj.TrailsData == nil {
		oj.TrailsData = NewTrailsJD()
	}
	if oj.AreasData == nil {
		oj.AreasData = NewAreasJD()
	}
	if oj.SessionsData == nil {
		oj.SessionsData = NewSessionsJD()
	}
}

func FromData(d *OreeJD) OreeJson {
	d.EnsureInitialized()
	return OreeJson{
		OreeJD: d,
	}
}

const digitMapping = "abcdefghijklmnopqrstuvwxyz0123456789"

func (o OreeJson) getAndIncId() string {
	id := o.NextId
	o.NextId++
	var b strings.Builder
	for {
		b.WriteByte(digitMapping[id%36])
		id /= 36
		if id == 0 {
			break
		}
		id--
	}
	return b.String()
}

func (o OreeJson) Trails() oree.TrailsI {
	return TrailsFromData(o.TrailsData, o)
}

func (o OreeJson) Areas() oree.AreasI {
	return AreasFromData(o.AreasData, o)
}

func (o OreeJson) Sessions() oree.SessionsI {
	return SessionsFromData(o.SessionsData, o)
}
