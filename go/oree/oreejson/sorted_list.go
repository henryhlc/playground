package oreejson

type ItemComparator[E any] interface {
	Compare(*E, *E) int
}

type SortedListOJ[I comparable, E any, D any, H any] struct {
	*ListJD[I, E]
	Converter  ItemHandleConverter[I, E, D, H]
	Comparator ItemComparator[E]
}

func SortedListFromData[I comparable, E any, D any, H any](
	d *ListJD[I, E], converter ItemHandleConverter[I, E, D, H], comparator ItemComparator[E]) SortedListOJ[I, E, D, H] {
	return SortedListOJ[I, E, D, H]{
		ListJD:     d,
		Converter:  converter,
		Comparator: comparator,
	}
}

func (sl SortedListOJ[I, E, D, H]) Len() int {
	return sl.ListJD.Len()
}

func (sl SortedListOJ[I, E, D, H]) Create(d D) H {
	item := sl.Converter.newItem(d)
	sl.PlaceItemFront(item)
	sl.FixOrder(item)
	return sl.Converter.itemToHandle(item)
}

func (sl SortedListOJ[I, E, D, H]) FixOrder(item ListItem[I, E]) {
	prevItems := sl.ListJD.NItemsBefore(1, item)

	shouldMoveBackwards := false
	for len(prevItems) > 0 && sl.Comparator.Compare(item.Elem, prevItems[0].Elem) < 0 {
		shouldMoveBackwards = true
		prevItems = sl.ListJD.NItemsBefore(1, prevItems[0])
	}
	if shouldMoveBackwards {
		if len(prevItems) == 0 {
			sl.ListJD.PlaceItemFront(item)
		} else {
			sl.ListJD.PlaceItemAfter(item, prevItems[0])
		}
	}

	nextItems := sl.ListJD.NItemsAfter(1, item)
	shouldMoveForward := false
	for len(nextItems) > 0 && sl.Comparator.Compare(item.Elem, nextItems[0].Elem) > 0 {
		shouldMoveForward = true
		nextItems = sl.ListJD.NItemsAfter(1, nextItems[0])
	}
	if shouldMoveForward {
		if len(nextItems) == 0 {
			sl.ListJD.PlaceItemBack(item)
		} else {
			sl.ListJD.PlaceItemBefore(item, nextItems[0])
		}
	}
}

func (sl SortedListOJ[I, E, D, H]) WithId(id I) (H, bool) {
	return ListFromData(sl.ListJD, sl.Converter).WithId(id)
}

func (sl SortedListOJ[I, E, D, H]) FirstN(n int) []H {
	return ListFromData(sl.ListJD, sl.Converter).FirstN(n)
}

func (sl SortedListOJ[I, E, D, H]) LastN(n int) []H {
	return ListFromData(sl.ListJD, sl.Converter).LastN(n)
}

func (sl SortedListOJ[I, E, D, H]) NAfter(n int, h H) []H {
	return ListFromData(sl.ListJD, sl.Converter).NAfter(n, h)
}

func (sl SortedListOJ[I, E, D, H]) NBefore(n int, h H) []H {
	return ListFromData(sl.ListJD, sl.Converter).NBefore(n, h)
}

func (sl SortedListOJ[I, E, D, H]) Delete(h H) {
	ListFromData(sl.ListJD, sl.Converter).Delete(h)
}
