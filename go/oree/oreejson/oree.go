package oreejson

import (
	"strings"

	"github.com/henryhlc/playground/go/oree"
)

type OreeJD struct {
	NextId                 int                   `json:"nextId"`
	BlocksData             *BlocksJD             `json:"blocks"`
	TrailsData             *TrailsJD             `json:"trails"`
	AreasData              *AreasJD              `json:"areas"`
	SessionsData           *SessionsJD           `json:"sessions"`
	OpenSessionManagerData *OpenSessionManagerJD `json:"openSession"`
}

type OreeJson struct {
	*OreeJD
}

func NewOreeJD() *OreeJD {
	return &OreeJD{
		NextId:       0,
		BlocksData:   NewBlocksJD(),
		TrailsData:   NewTrailsJD(),
		AreasData:    NewAreasJD(),
		SessionsData: NewSessionsJD(),
	}
}

func (oj *OreeJD) EnsureInitialized() {
	if oj.BlocksData == nil {
		oj.BlocksData = NewBlocksJD()
	}
	if oj.TrailsData == nil {
		oj.TrailsData = NewTrailsJD()
	}
	if oj.AreasData == nil {
		oj.AreasData = NewAreasJD()
	}
	if oj.SessionsData == nil {
		oj.SessionsData = NewSessionsJD()
	}
	if oj.OpenSessionManagerData == nil {
		oj.OpenSessionManagerData = NewOpenSessionManagerJD()
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

func (o OreeJson) Blocks() oree.BlocksI {
	return BlocksFromData(o.BlocksData, o)
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

func (o OreeJson) OpenSessionManager() oree.OpenSessionManagerI {
	return OpenSessionManagerFromData(o.OpenSessionManagerData, o)
}
