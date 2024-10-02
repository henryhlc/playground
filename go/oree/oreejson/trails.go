package oreejson

import (
	"github.com/henryhlc/playground/go/oree"
)

type TrailsJsonData struct {
	*OrderedListJsonData[oree.TrailId, TrailJsonData]
}

func NewTrailsJsonData() *TrailsJsonData {
	return &TrailsJsonData{
		OrderedListJsonData: NewOrderedListJsonData[oree.TrailId, TrailJsonData](),
	}
}

type TrailsJson struct {
	OrderedListJson[oree.TrailId, TrailJsonData, oree.Trail, oree.TrailI]
	oreeJson OreeJson
}

func TrailsFromData(data *TrailsJsonData, oreeJson OreeJson) TrailsJson {
	return TrailsJson{
		OrderedListJson: OrderedListJsonFromData(
			data.OrderedListJsonData,
			newItemTrailIConverter(oreeJson),
		),
		oreeJson: oreeJson,
	}
}

type ItemTrailIConverter struct {
	oreeJson OreeJson
}

func newItemTrailIConverter(oreeJson OreeJson) ItemTrailIConverter {
	return ItemTrailIConverter{oreeJson: oreeJson}
}

func (c ItemTrailIConverter) emptyHandle() oree.TrailI {
	return TrailJson{}
}

func (c ItemTrailIConverter) newItem(d oree.Trail) ListItem[oree.TrailId, TrailJsonData] {
	id := oree.TrailId(c.oreeJson.getAndIncId())
	return ListItem[oree.TrailId, TrailJsonData]{
		Id:   id,
		Elem: NewTrailJsonData(d),
	}
}

func (c ItemTrailIConverter) updatedItem(
	item ListItem[oree.TrailId, TrailJsonData],
	d oree.Trail) ListItem[oree.TrailId, TrailJsonData] {
	return ListItem[oree.TrailId, TrailJsonData]{
		Id:   item.Id,
		Elem: NewTrailJsonData(d),
	}
}

func (c ItemTrailIConverter) itemToHandle(
	item ListItem[oree.TrailId, TrailJsonData]) oree.TrailI {
	return TrailJson{
		TrailJsonData: item.Elem,
		oreeJson:      c.oreeJson,
		id:            item.Id,
	}
}

func (c ItemTrailIConverter) handleToItem(h oree.TrailI) ListItem[oree.TrailId, TrailJsonData] {
	tj := h.(TrailJson)
	return ListItem[oree.TrailId, TrailJsonData]{
		Id:   tj.id,
		Elem: tj.TrailJsonData,
	}
}
