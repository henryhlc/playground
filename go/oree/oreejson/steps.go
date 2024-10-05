package oreejson

import "github.com/henryhlc/playground/go/oree"

type StepsJD struct {
	*OrderedListJD[oree.StepId, StepJD]
}

type StepsOJ struct {
	OrderedListOJ[oree.StepId, StepJD, oree.Step, oree.StepI]
	oreeJson OreeJson
}

func NewStepsJD() *StepsJD {
	return &StepsJD{
		OrderedListJD: NewOrderedListJD[oree.StepId, StepJD](),
	}
}

func (s *StepsJD) EnsureInitialized() {
	if s.OrderedListJD == nil {
		s.OrderedListJD = NewOrderedListJD[oree.StepId, StepJD]()
	}
}

func StepsFromData(data *StepsJD, oj OreeJson) StepsOJ {
	data.EnsureInitialized()
	return StepsOJ{
		OrderedListOJ: OrderedListFromData(
			data.OrderedListJD,
			newItemStepIConverter(oj),
		),
		oreeJson: oj,
	}
}

type ItemStepIConverter struct {
	oreeJson OreeJson
}

func newItemStepIConverter(oj OreeJson) ItemStepIConverter {
	return ItemStepIConverter{oreeJson: oj}
}

func (c ItemStepIConverter) emptyHandle() oree.StepI {
	return StepOJ{}
}

func (c ItemStepIConverter) newItem(d oree.Step) ListItem[oree.StepId, StepJD] {
	id := oree.StepId(c.oreeJson.getAndIncId())
	return ListItem[oree.StepId, StepJD]{
		Id:   id,
		Elem: NewStepJD(d),
	}
}

func (c ItemStepIConverter) updatedItem(
	item ListItem[oree.StepId, StepJD],
	d oree.Step) ListItem[oree.StepId, StepJD] {
	return ListItem[oree.StepId, StepJD]{
		Id:   item.Id,
		Elem: NewStepJD(d),
	}
}

func (c ItemStepIConverter) itemToHandle(
	item ListItem[oree.StepId, StepJD]) oree.StepI {
	return StepOJ{
		StepJD:   item.Elem,
		oreeJson: c.oreeJson,
		id:       item.Id,
	}
}

func (c ItemStepIConverter) handleToItem(h oree.StepI) ListItem[oree.StepId, StepJD] {
	sj := h.(StepOJ)
	return ListItem[oree.StepId, StepJD]{
		Id:   sj.id,
		Elem: sj.StepJD,
	}
}
