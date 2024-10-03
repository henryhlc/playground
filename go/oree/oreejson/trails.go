package oreejson

import (
	"github.com/henryhlc/playground/go/oree"
)

type TrailsJD struct {
	*OrderedListJD[oree.TrailId, TrailJD]
}

type TrailsOJ struct {
	OrderedListOJ[oree.TrailId, TrailJD, oree.Trail, oree.TrailI]
	oreeJson OreeJson
}

func NewTrailsJD() *TrailsJD {
	return &TrailsJD{
		OrderedListJD: NewOrderedListJD[oree.TrailId, TrailJD](),
	}
}

func TrailsFromData(data *TrailsJD, oj OreeJson) TrailsOJ {
	return TrailsOJ{
		OrderedListOJ: OrderedListFromData(
			data.OrderedListJD,
			newItemTrailIConverter(oj),
		),
		oreeJson: oj,
	}
}

type ItemTrailIConverter struct {
	oreeJson OreeJson
}

func newItemTrailIConverter(oj OreeJson) ItemTrailIConverter {
	return ItemTrailIConverter{oreeJson: oj}
}

func (c ItemTrailIConverter) emptyHandle() oree.TrailI {
	return TrailOJ{}
}

func (c ItemTrailIConverter) newItem(d oree.Trail) ListItem[oree.TrailId, TrailJD] {
	id := oree.TrailId(c.oreeJson.getAndIncId())
	return ListItem[oree.TrailId, TrailJD]{
		Id:   id,
		Elem: NewTrailJD(d),
	}
}

func (c ItemTrailIConverter) updatedItem(
	item ListItem[oree.TrailId, TrailJD],
	d oree.Trail) ListItem[oree.TrailId, TrailJD] {
	return ListItem[oree.TrailId, TrailJD]{
		Id:   item.Id,
		Elem: NewTrailJD(d),
	}
}

func (c ItemTrailIConverter) itemToHandle(
	item ListItem[oree.TrailId, TrailJD]) oree.TrailI {
	return TrailOJ{
		TrailJD:  item.Elem,
		oreeJson: c.oreeJson,
		id:       item.Id,
	}
}

func (c ItemTrailIConverter) handleToItem(h oree.TrailI) ListItem[oree.TrailId, TrailJD] {
	tj := h.(TrailOJ)
	return ListItem[oree.TrailId, TrailJD]{
		Id:   tj.id,
		Elem: tj.TrailJD,
	}
}
