package oreejson

import "github.com/henryhlc/playground/go/oree"

type AreasJD struct {
	*ListJD[oree.AreaId, AreaJD]
}

type AreasOJ struct {
	ListOJ[oree.AreaId, AreaJD, oree.Area, oree.AreaI]
}

func NewAreasJD() *AreasJD {
	return &AreasJD{
		ListJD: NewListJD[oree.AreaId, AreaJD](),
	}
}

func AreasFromData(d *AreasJD, oj OreeJson) AreasOJ {
	return AreasOJ{
		ListOJ: ListFromData(d.ListJD, newItemAreaConverter(oj)),
	}
}

type ItemAreaIConverter struct {
	oreeJson OreeJson
}

func newItemAreaConverter(oj OreeJson) ItemAreaIConverter {
	return ItemAreaIConverter{
		oreeJson: oj,
	}
}

func (c ItemAreaIConverter) emptyHandle() oree.AreaI {
	return AreaOJ{}
}

func (c ItemAreaIConverter) newItem(d oree.Area) ListItem[oree.AreaId, AreaJD] {
	id := oree.AreaId(c.oreeJson.getAndIncId())
	return ListItem[oree.AreaId, AreaJD]{
		Id:   id,
		Elem: NewAreaJD(d),
	}
}

func (c ItemAreaIConverter) updatedItem(
	item ListItem[oree.AreaId, AreaJD],
	d oree.Area) ListItem[oree.AreaId, AreaJD] {
	return ListItem[oree.AreaId, AreaJD]{
		Id:   item.Id,
		Elem: NewAreaJD(d),
	}
}

func (c ItemAreaIConverter) itemToHandle(
	item ListItem[oree.AreaId, AreaJD]) oree.AreaI {
	return AreaFromData(item.Elem, c.oreeJson, item.Id)
}

func (c ItemAreaIConverter) handleToItem(h oree.AreaI) ListItem[oree.AreaId, AreaJD] {
	a := h.(AreaOJ)
	return ListItem[oree.AreaId, AreaJD]{
		Id:   a.id,
		Elem: a.AreaJD,
	}
}
