package oreejson

import (
	"github.com/henryhlc/playground/go/oree"
)

type TrailsJsonData struct {
	*OrderedListJson[oree.TrailId, TrailJsonData]
}

type TrailsJson struct {
	*TrailsJsonData
	oreeJson OreeJson
}

func NewTrailsJsonData() *TrailsJsonData {
	return &TrailsJsonData{
		NewOrderedListJson[oree.TrailId, TrailJsonData](),
	}
}

func (t TrailsJson) newTrailItem(trail oree.Trail) ListItem[oree.TrailId, TrailJsonData] {
	id := oree.TrailId(t.oreeJson.getAndIncId())
	data := NewTrailJsonData(trail)
	return ListItem[oree.TrailId, TrailJsonData]{
		Id:   id,
		Elem: data,
	}
}

func (t TrailsJson) createTrailIFromItem(item ListItem[oree.TrailId, TrailJsonData]) oree.TrailI {
	return TrailJson{
		TrailJsonData: item.Elem,
		oreeJson:      t.oreeJson,
		id:            item.Id,
	}
}

func (t TrailsJson) createItemFromTrailJson(tj TrailJson) ListItem[oree.TrailId, TrailJsonData] {
	return ListItem[oree.TrailId, TrailJsonData]{
		Id:   tj.id,
		Elem: tj.TrailJsonData,
	}
}

func (t TrailsJson) createTrailIsFromItems(items []ListItem[oree.TrailId, TrailJsonData]) []oree.TrailI {
	tjs := make([]oree.TrailI, len(items))
	for i, item := range items {
		tjs[i] = t.createTrailIFromItem(item)
	}
	return tjs
}

func (t TrailsJson) Len() int {
	return t.TrailsJsonData.Len()
}

func (t TrailsJson) CreateFront(trail oree.Trail) oree.TrailI {
	item := t.newTrailItem(trail)
	t.PlaceItemFront(item)
	return t.createTrailIFromItem(item)
}

func (t TrailsJson) CreateBack(trail oree.Trail) oree.TrailI {
	item := t.newTrailItem(trail)
	t.PlaceItemBack(item)
	return t.createTrailIFromItem(item)
}

func (t TrailsJson) CreateBefore(trail oree.Trail, nbr oree.TrailI) oree.TrailI {
	item := t.newTrailItem(trail)
	nbrItem, _ := t.ItemWithId(nbr.Id())
	t.PlaceItemBefore(item, nbrItem)
	return t.createTrailIFromItem(item)
}

func (t TrailsJson) CreateAfter(trail oree.Trail, nbr oree.TrailI) oree.TrailI {
	item := t.newTrailItem(trail)
	nbrItem, _ := t.ItemWithId(nbr.Id())
	t.PlaceItemAfter(item, nbrItem)
	return t.createTrailIFromItem(item)
}

func (t TrailsJson) WithId(id oree.TrailId) (tj oree.TrailI, ok bool) {
	item, ok := t.ItemWithId(id)
	if !ok {
		return TrailJson{}, false
	}
	return t.createTrailIFromItem(item), true
}

func (t TrailsJson) FirstN(n int) []oree.TrailI {
	return t.createTrailIsFromItems(t.FirstNItems(n))
}

func (t TrailsJson) LastN(n int) []oree.TrailI {
	return t.createTrailIsFromItems(t.LastNItems(n))
}

func (t TrailsJson) NAfter(n int, ti oree.TrailI) []oree.TrailI {
	tj := ti.(TrailJson)
	return t.createTrailIsFromItems(
		t.NItemsAfter(n, t.createItemFromTrailJson(tj)),
	)
}

func (t TrailsJson) NBefore(n int, ti oree.TrailI) []oree.TrailI {
	tj := ti.(TrailJson)
	return t.createTrailIsFromItems(
		t.NItemsBefore(n, t.createItemFromTrailJson(tj)),
	)
}

func (t TrailsJson) PlaceFront(ti oree.TrailI) {
	tj := ti.(TrailJson)
	t.PlaceItemFront(t.createItemFromTrailJson(tj))
}

func (t TrailsJson) PlaceBack(ti oree.TrailI) {
	tj := ti.(TrailJson)
	t.PlaceItemBack(t.createItemFromTrailJson(tj))
}

func (t TrailsJson) PlaceBefore(ti, nbri oree.TrailI) {
	tj := ti.(TrailJson)
	nbr := nbri.(TrailJson)
	t.PlaceItemBefore(
		t.createItemFromTrailJson(tj),
		t.createItemFromTrailJson(nbr),
	)
}

func (t TrailsJson) PlaceAfter(ti, nbri oree.TrailI) {
	tj := ti.(TrailJson)
	nbr := nbri.(TrailJson)
	t.PlaceItemAfter(
		t.createItemFromTrailJson(tj),
		t.createItemFromTrailJson(nbr),
	)
}

func (t TrailsJson) Update(ti oree.TrailI, trail oree.Trail) {
	tj := ti.(TrailJson)
	t.UpdateItem(ListItem[oree.TrailId, TrailJsonData]{
		Id:   tj.id,
		Elem: NewTrailJsonData(trail),
	})
}

func (t TrailsJson) Delete(ti oree.TrailI) {
	tj := ti.(TrailJson)
	t.DeleteItem(t.createItemFromTrailJson(tj))
}
