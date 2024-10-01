package oreejson

import "github.com/henryhlc/playground/go/oree"

type TrailsJsonData struct {
	ById map[oree.TrailId]*TrailJsonData `json:"byId"`
}

func NewTrailsJsonData() *TrailsJsonData {
	return &TrailsJsonData{
		ById: map[oree.TrailId]*TrailJsonData{},
	}
}

type TrailJsonData struct {
	Description string `json:"description"`
}

type TrailsJson struct {
	*TrailsJsonData
	oreeJson OreeJson
}

func (t TrailsJson) CreateTrail(trail oree.Trail) oree.TrailId {
	id := oree.TrailId(t.oreeJson.getAndIncId())
	t.ById[id] = &TrailJsonData{
		Description: trail.Description,
	}
	return id
}

func (t TrailsJson) TrailWithId(id oree.TrailId) oree.TrailI {
	return TrailJson{
		TrailJsonData: t.ById[id],
		id:            id,
		oreeJson:      t.oreeJson,
	}
}

type TrailJson struct {
	*TrailJsonData
	oreeJson OreeJson
	id       oree.TrailId
}

func (t TrailJson) Data() oree.Trail {
	return oree.Trail{
		Description: t.Description,
	}
}
