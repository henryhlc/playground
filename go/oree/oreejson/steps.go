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

func StepsFromData(data *StepsJD, oj OreeJson, t TrailOJ, status oree.StepStatus) StepsOJ {
	data.EnsureInitialized()
	return StepsOJ{
		OrderedListOJ: OrderedListFromData(
			data.OrderedListJD,
			newItemStepIConverter(oj, t, status),
		),
		oreeJson: oj,
	}
}

type ItemStepIConverter struct {
	oreeJson OreeJson
	status   oree.StepStatus
	trail    TrailOJ
}

func newItemStepIConverter(oj OreeJson, t TrailOJ, s oree.StepStatus) ItemStepIConverter {
	return ItemStepIConverter{oreeJson: oj, trail: t, status: s}
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
	return StepFromData(item.Elem, c.oreeJson, c.trail, item.Id, c.status)
}

func (c ItemStepIConverter) handleToItem(h oree.StepI) ListItem[oree.StepId, StepJD] {
	sj := h.(StepOJ)
	return ListItem[oree.StepId, StepJD]{
		Id:   sj.id,
		Elem: sj.StepJD,
	}
}
