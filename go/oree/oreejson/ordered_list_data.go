package oreejson

import (
	"slices"
)

type OrderedListJD[I comparable, E any] struct {
	ById map[I]*E `json:"byId"`
	Next map[I]I  `json:"next"`
	Prev map[I]I  `json:"prev"`
	Head I        `json:"head"`
	Tail I        `json:"tail"`
}

type ListItem[I comparable, E any] struct {
	Id   I
	Elem *E
}

func NewOrderedListJD[I comparable, E any]() *OrderedListJD[I, E] {
	return &OrderedListJD[I, E]{
		ById: map[I]*E{},
		Next: map[I]I{},
		Prev: map[I]I{},
	}
}

func (ol OrderedListJD[I, E]) Len() int {
	return len(ol.ById)
}

func (ol *OrderedListJD[I, E]) PlaceItemBack(item ListItem[I, E]) {
	if item, ok := ol.ItemWithId(item.Id); ok {
		ol.DeleteItem(item)
	}
	if ol.Len() == 0 {
		ol.Head = item.Id
		ol.Tail = item.Id
	} else {
		ol.Next[ol.Tail] = item.Id
		ol.Prev[item.Id] = ol.Tail
		ol.Tail = item.Id
	}
	ol.ById[item.Id] = item.Elem
}

func (ol *OrderedListJD[I, E]) PlaceItemFront(item ListItem[I, E]) {
	if item, ok := ol.ItemWithId(item.Id); ok {
		ol.DeleteItem(item)
	}
	if ol.Len() == 0 {
		ol.PlaceItemBack(item)
		return
	}
	ol.Prev[ol.Head] = item.Id
	ol.Next[item.Id] = ol.Head
	ol.Head = item.Id
	ol.ById[item.Id] = item.Elem
}

// Requires nbr to be an item in the list.
func (ol *OrderedListJD[I, E]) PlaceItemBefore(item, nbr ListItem[I, E]) {
	if item, ok := ol.ItemWithId(item.Id); ok {
		ol.DeleteItem(item)
	}
	if ol.Len() == 0 || nbr.Id == ol.Head {
		ol.PlaceItemFront(item)
		return
	}
	ol.ById[item.Id] = item.Elem
	nbrPrevId := ol.Prev[nbr.Id]
	ol.Next[nbrPrevId] = item.Id
	ol.Prev[nbr.Id] = item.Id
	ol.Next[item.Id] = nbr.Id
	ol.Prev[item.Id] = nbrPrevId
}

// Requires nbr to be an item in the list.
func (ol *OrderedListJD[I, E]) PlaceItemAfter(item, nbr ListItem[I, E]) {
	if item, ok := ol.ItemWithId(item.Id); ok {
		ol.DeleteItem(item)
	}
	if ol.Len() == 0 || nbr.Id == ol.Tail {
		ol.PlaceItemBack(item)
		return
	}
	ol.ById[item.Id] = item.Elem
	nbrNextId := ol.Next[nbr.Id]
	ol.Next[nbr.Id] = item.Id
	ol.Prev[nbrNextId] = item.Id
	ol.Next[item.Id] = nbrNextId
	ol.Prev[item.Id] = nbr.Id
}

func (ol *OrderedListJD[I, E]) ItemWithId(id I) (item ListItem[I, E], ok bool) {
	e, ok := ol.ById[id]
	if !ok {
		return ListItem[I, E]{}, false
	}
	return ListItem[I, E]{
		Id:   id,
		Elem: e,
	}, true
}

func (ol *OrderedListJD[I, E]) FirstNItems(n int) []ListItem[I, E] {
	if ol.Len() == 0 {
		return nil
	}
	headItem, _ := ol.ItemWithId(ol.Head)
	items := []ListItem[I, E]{headItem}
	if ol.Len() == 1 || n == 1 {
		return items
	}
	return append(items, ol.NItemsAfter(n-1, headItem)...)
}

func (ol *OrderedListJD[I, E]) NItemsAfter(n int, item ListItem[I, E]) []ListItem[I, E] {
	currId, ok := ol.Next[item.Id]
	if !ok {
		return nil
	}
	items := []ListItem[I, E]{}
	for range n {
		currItem, _ := ol.ItemWithId(currId)
		items = append(items, currItem)
		if nextId, ok := ol.Next[currId]; ok {
			currId = nextId
		} else {
			break
		}
	}
	return items
}

func (ol *OrderedListJD[I, E]) LastNItems(n int) []ListItem[I, E] {
	if ol.Len() == 0 {
		return nil
	}
	tailItem, _ := ol.ItemWithId(ol.Tail)
	if ol.Len() == 1 || n == 1 {
		return []ListItem[I, E]{tailItem}
	}
	return append(ol.NItemsBefore(n-1, tailItem), tailItem)

}

func (ol *OrderedListJD[I, E]) NItemsBefore(n int, item ListItem[I, E]) []ListItem[I, E] {
	currId, ok := ol.Prev[item.Id]
	if !ok {
		return nil
	}
	items := []ListItem[I, E]{}
	for range n {
		currItem, _ := ol.ItemWithId(currId)
		items = append(items, currItem)
		if prevId, ok := ol.Prev[currId]; ok {
			currId = prevId
		} else {
			break
		}
	}
	slices.Reverse(items)
	return items
}

func (ol *OrderedListJD[I, E]) DeleteItem(item ListItem[I, E]) {
	if ol.Len() == 0 {
		return
	}
	item, ok := ol.ItemWithId(item.Id)
	if !ok {
		return
	}

	switch {
	case item.Id == ol.Head && item.Id == ol.Tail:
		break
	case item.Id == ol.Head:
		nextId := ol.Next[item.Id]
		ol.Head = nextId
		delete(ol.Prev, nextId)
		delete(ol.Next, item.Id)
	case item.Id == ol.Tail:
		prevId := ol.Prev[item.Id]
		ol.Tail = prevId
		delete(ol.Prev, item.Id)
		delete(ol.Next, prevId)
	default:
		prevId, nextId := ol.Prev[item.Id], ol.Next[item.Id]
		ol.Next[prevId] = nextId
		ol.Prev[nextId] = prevId
		delete(ol.Prev, item.Id)
		delete(ol.Next, item.Id)
	}
	delete(ol.ById, item.Id)
}
