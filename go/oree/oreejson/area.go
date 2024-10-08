package oreejson

import "github.com/henryhlc/playground/go/oree"

type AreaJD struct {
	Description string `json:"description"`
	TrailIds    *RefListJD[oree.TrailId]
}

type AreaOJ struct {
	*AreaJD
	id        oree.AreaId
	trailRefs RefListOJ[oree.TrailId, oree.TrailI]
}

func NewAreaJD(a oree.Area) *AreaJD {
	return &AreaJD{
		Description: a.Description,
		TrailIds:    NewRefListJD[oree.TrailId](),
	}
}

func AreaFromData(d *AreaJD, oj OreeJson, id oree.AreaId) AreaOJ {
	return AreaOJ{
		AreaJD:    d,
		id:        id,
		trailRefs: RefListFromData(d.TrailIds, TrailResolver{oreeJson: oj}, oree.TrailId("")),
	}
}

type TrailResolver struct {
	oreeJson OreeJson
}

func (tr TrailResolver) ResolveRef(id oree.TrailId) (oree.TrailI, bool) {
	return tr.oreeJson.Trails().WithId(id)
}

func (tr TrailResolver) ExtractRef(t oree.TrailI) oree.TrailId {
	return t.Id()
}

func (tr TrailResolver) EmptyRef() oree.TrailId {
	return oree.TrailId("")
}

func (a AreaOJ) Id() oree.AreaId {
	return a.id
}

func (a AreaOJ) Data() oree.Area {
	return oree.Area{
		Description: a.Description,
	}
}

func (a AreaOJ) Update(area oree.Area) {
	a.AreaJD.Description = area.Description
}

func (a AreaOJ) Trails() oree.RefListI[oree.TrailI, oree.TrailId] {
	return a.trailRefs
}
