package oreejson

import "github.com/henryhlc/playground/go/oree"

type SessionsJD struct {
	*ListJD[oree.SessionId, SessionJD]
}

type SessionsOJ struct {
	SortedListOJ[oree.SessionId, SessionJD, oree.Session, oree.SessionI]
	oreeJson OreeJson
}

func NewSessionsJD() *SessionsJD {
	return &SessionsJD{
		ListJD: NewListJD[oree.SessionId, SessionJD](),
	}
}

func SessionsFromData(d *SessionsJD, oj OreeJson) SessionsOJ {
	return SessionsOJ{
		SortedListOJ: SortedListFromData(
			d.ListJD,
			newItemSessionConverter(oj),
			SessionComparator{},
		),
		oreeJson: oj,
	}
}

type ItemSessionConverter struct {
	oreeJson OreeJson
}

func newItemSessionConverter(oj OreeJson) ItemSessionConverter {
	return ItemSessionConverter{oreeJson: oj}
}

func (c ItemSessionConverter) emptyHandle() oree.SessionI {
	return SessionOJ{}
}

func (c ItemSessionConverter) newItem(d oree.Session) ListItem[oree.SessionId, SessionJD] {
	return ListItem[oree.SessionId, SessionJD]{
		Id:   oree.SessionId(c.oreeJson.getAndIncId()),
		Elem: NewSessionJD(d),
	}
}

func (c ItemSessionConverter) itemToHandle(item ListItem[oree.SessionId, SessionJD]) oree.SessionI {
	return SessionFromData(item.Elem, c.oreeJson, item.Id)
}

func (c ItemSessionConverter) handleToItem(s oree.SessionI) ListItem[oree.SessionId, SessionJD] {
	soj := s.(SessionOJ)
	return ListItem[oree.SessionId, SessionJD]{
		Id:   s.Id(),
		Elem: soj.SessionJD,
	}
}

type SessionComparator struct{}

func (SessionComparator) Compare(a, b *SessionJD) int {
	return a.StartTime.Compare(b.StartTime)
}
